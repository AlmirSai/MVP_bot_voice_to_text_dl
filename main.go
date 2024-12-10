package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var selectedModel = "tiny" // Модель по умолчанию

func main() {
	// Проверяем наличие необходимых утилит
	if err := checkDependencies(); err != nil {
		log.Fatalf("Missing dependencies: %v", err)
	}

	// Заменить на свой токен
	bot, err := tgbotapi.NewBotAPI("TELEGRAM_TOKEN")
	if err != nil {
		log.Fatalf("Error initializing bot: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Если сообщение — текстовое и содержит команду выбора модели
		if update.Message.Text != "" {
			handleTextMessage(bot, update.Message)
			continue
		}

		// Если сообщение — голосовое
		if update.Message.Voice != nil {
			handleVoiceMessage(bot, update.Message)
		}
	}
}

func checkDependencies() error {
	requiredCommands := []string{"ffmpeg", "whisper"}
	for _, cmd := range requiredCommands {
		if _, err := exec.LookPath(cmd); err != nil {
			return fmt.Errorf("required command not found: %s", cmd)
		}
	}
	return nil
}

func handleTextMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	text := strings.ToLower(message.Text)

	if strings.Contains(text, "model:") {
		parts := strings.Split(text, ":")
		if len(parts) == 2 {
			model := strings.TrimSpace(parts[1])
			if model == "tiny" || model == "large" {
				selectedModel = model
				response := fmt.Sprintf("Model successfully changed to: %s", selectedModel)
				msg := tgbotapi.NewMessage(message.Chat.ID, response)
				bot.Send(msg)
				return
			}
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "Invalid model! Available options: tiny, large.")
		bot.Send(msg)
	}
}

func handleVoiceMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	voice := message.Voice
	fileConfig := tgbotapi.FileConfig{FileID: voice.FileID}
	file, err := bot.GetFile(fileConfig)
	if err != nil {
		log.Printf("Failed to get voice file: %v", err)
		return
	}

	downloadURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath)
	log.Printf("Voice file URL: %s", downloadURL)

	localPath := "voice.ogg"
	err = downloadFile(localPath, downloadURL)
	if err != nil {
		log.Printf("Failed to download voice message: %v", err)
		return
	}
	defer os.Remove(localPath)

	transcription, err := transcribeWithWhisper(localPath, selectedModel)
	if err != nil {
		log.Printf("Failed to transcribe voice message: %v", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Failed to transcribe the voice message. Please try again.")
		bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, transcription)
	bot.Send(msg)
}

func downloadFile(filepath, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func transcribeWithWhisper(audioPath, model string) (string, error) {
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
