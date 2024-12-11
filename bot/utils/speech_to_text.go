package utils

import (
	"fmt"
	"os"
	"os/exec"

	"tg-whisper-bot/bot/utils/logger"
)

func TranscribeWithWhisper(audioPath, model string) (string, error) {
	logFilePath := "storage/logs/speech_to_text.log"

	if _, err := os.Stat("storage/logs"); os.IsNotExist(err) {
		err := os.MkdirAll("storage/logs", 0755)
		if err != nil {
			fmt.Println("Failed to create log directory:", err)
			return "", err
		}
	}

	logInstance, err := logger.GetLogger(logFilePath)
	if err != nil {
		fmt.Println("Error initializing logger:", err)
		return "", err
	}
	defer logInstance.Close()

	logInstance.Info(fmt.Sprintf("Started transcribing audio file: %s with model: %s", audioPath, model))

	wavPath := "voice.wav"
	cmdConvert := exec.Command("ffmpeg", "-i", audioPath, "-ar", "16000", "-ac", "1", wavPath)
	output, err := cmdConvert.CombinedOutput()
	logInstance.Info(fmt.Sprintf("FFmpeg output: %s", string(output)))
	if err != nil {
		logInstance.Error(fmt.Sprintf("Error converting audio file: %v", err))
		return "", fmt.Errorf("error converting audio file: %v", err)
	}
	defer os.Remove(wavPath)

	outputPath := "voice.txt"
	cmd := exec.Command("whisper", wavPath, "--language", "ru", "--model", model, "--output_format", "txt")
	whisperOutput, err := cmd.CombinedOutput()
	logInstance.Info(fmt.Sprintf("Whisper output: %s", string(whisperOutput)))
	if err != nil {
		logInstance.Error(fmt.Sprintf("Error running Whisper: %v", err))
		return "", fmt.Errorf("error running Whisper: %v", err)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		logInstance.Error("Transcription file not found.")
		return "", fmt.Errorf("transcription file not found")
	}
	defer os.Remove(outputPath)

	transcription, err := os.ReadFile(outputPath)
	if err != nil {
		logInstance.Error(fmt.Sprintf("Error reading transcription file: %v", err))
		return "", fmt.Errorf("error reading transcription file: %v", err)
	}

	logInstance.Info("Transcription completed successfully.")
	return string(transcription), nil
}
