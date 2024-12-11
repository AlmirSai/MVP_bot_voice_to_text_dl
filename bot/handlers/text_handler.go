package handlers

import (
	"fmt"
	"strings"

	"tg-whisper-bot/bot/utils/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleTextMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	text := strings.ToLower(message.Text)

	logFilePath := "storage/logs/text_handler.log"

	logInstance, err := logger.GetLogger(logFilePath)
	if err != nil {
		fmt.Println("Error initializing logger:", err)
		return
	}
	defer logInstance.Close()

	logInstance.Info(fmt.Sprintf("Received message: %s", text))

	if text == "/model" {
		HandleModelCommand(bot, message, logInstance)
		return
	}

	logInstance.Info("No relevant command detected.")
}
