// bot/handlers/text_handler.go
package handlers

import (
	"fmt"
	"strings"

	"tg-whisper-bot/bot/utils/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleTextMessage processes a text message from the Telegram bot.
// It handles the following commands: /model, /info.
// For unrecognized commands, it sends a default response.
func HandleTextMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	text := strings.TrimSpace(strings.ToLower(message.Text))
	
	logFilePath := "storage/logs/text_handler.log"

	logInstance, err := logger.GetLogger(logFilePath)
	if err != nil {
		fmt.Printf("Error initialing logger: %v\n", err)
		return
	}
	defer logInstance.Close()

	logInstance.Info(fmt.Sprintf("Received message from user %s (ID: %d): %s", message.From.UserName, message.From.ID, message.Text))


	switch text {
	case "/model":
		HandleModelCommand(bot, message, logInstance)
	case "/info":
		HandleInfoCommand(bot, message, logInstance)
	default:
		logInstance.Info(fmt.Sprintf("Unrecognized command: %s", text))

		defaultResponse := "Sorry, I didn't understand that command. Use /info to learn more about my features or /model to choose a transcription model."
		msg := tgbotapi.NewMessage(message.Chat.ID, defaultResponse)
		if _, err := bot.Send(msg); err != nil {
			logInstance.Error(fmt.Sprintf("Error sending unrecognized command response: %v", err))
		}
	}
}
