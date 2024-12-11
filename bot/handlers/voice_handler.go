package handlers

import (
	"fmt"	
	"os"
	"tg-whisper-bot/bot/utils"
	"tg-whisper-bot/bot/utils/logger" 

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleVoiceMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	logInstance, err := logger.GetLogger("storage/logs/voice_handler.log")
	if err != nil {
		fmt.Println("Error initializing logger:", err)
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

	// TODO: change standart path
	localPath := "voice.ogg"

	err = utils.DownloadFile(localPath, downloadURL)
	if err != nil {
		logInstance.Error(fmt.Sprintf("Failed to download voice message: %v", err))
		return
	}
	defer os.Remove(localPath)

	transcription, err := utils.TranscribeWithWhisper(localPath, SelectedModel)
	if err != nil {
		logInstance.Error(fmt.Sprintf("Failed to transcribe voice message: %v", err))
		msg := tgbotapi.NewMessage(message.Chat.ID, "Failed to transcribe the voice message. Please try again.")
		bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, transcription)
	bot.Send(msg)
}
