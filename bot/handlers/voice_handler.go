// bot/handlers/voice_handler.go
package handlers

import (
	"fmt"
	"os"

	"tg-whisper-bot/bot/utils"
	"tg-whisper-bot/bot/utils/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleVoiceMessage processes a voice message from the Telegram bot.
func HandleVoiceMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	logInstance, err := logger.GetLogger("storage/logs/voice_handler.log")
	if err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		return
	}
	defer logInstance.Close()

	voice := message.Voice
	fileConfig := tgbotapi.FileConfig{FileID: voice.FileID}
	file, err := bot.GetFile(fileConfig)
	if err != nil {
		logInstance.Error(fmt.Sprintf("Failed to get voice file: %v", err))
		return
	}

	downloadURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath)
	logInstance.Info(fmt.Sprintf("Voice file URL: %s", downloadURL))

	// TODO: change standard path
	const localPath = "voice.ogg"

	if err := utils.DownloadFile(localPath, downloadURL); err != nil {
		logInstance.Error(fmt.Sprintf("Failed to download voice message: %v", err))
		return
	}
	defer func() {
		if err := os.Remove(localPath); err != nil {
			logInstance.Error(fmt.Sprintf("Failed to remove local file: %v", err))
		}
	}()

	transcription, err := utils.TranscribeWithWhisper(localPath, SelectedModel)
	if err != nil {
		logInstance.Error(fmt.Sprintf("Failed to transcribe voice message: %v", err))
		msg := tgbotapi.NewMessage(message.Chat.ID, "Failed to transcribe the voice message. Please try again.")
		if _, err := bot.Send(msg); err != nil {
			logInstance.Error(fmt.Sprintf("Failed to send message: %v", err))
		}
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, transcription)
	if _, err := bot.Send(msg); err != nil {
		logInstance.Error(fmt.Sprintf("Failed to send message: %v", err))
	}
}
