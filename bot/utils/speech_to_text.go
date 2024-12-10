package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func TranscribeWithWhisper(audioPath, model string) (string, error) {
	// TODO: change standart path
	wavPath := "voice.wav"
	cmdConvert := exec.Command("ffmpeg", "-i", audioPath, "-ar", "16000", "-ac", "1", wavPath)
	output, err := cmdConvert.CombinedOutput()
	log.Printf("FFmpeg output: %s", string(output))
	if err != nil {
		return "", fmt.Errorf("error converting audio file: %v", err)
	}
	defer os.Remove(wavPath)

	outputPath := "voice.txt"
	cmd := exec.Command("whisper", wavPath, "--language", "ru", "--model", model, "--output_format", "txt")
	whisperOutput, err := cmd.CombinedOutput()
	log.Printf("Whisper output: %s", string(whisperOutput))
	if err != nil {
		return "", fmt.Errorf("error running Whisper: %v", err)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return "", fmt.Errorf("transcription file not found")
	}
	defer os.Remove(outputPath)

	transcription, err := os.ReadFile(outputPath)
	if err != nil {
		return "", fmt.Errorf("error reading transcription file: %v", err)
	}

	return string(transcription), nil
}
