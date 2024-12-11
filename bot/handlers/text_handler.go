package handlers

import (
	"fmt"
	"os"
	"strings"

	"tg-whisper-bot/bot/utils/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var SelectedModel = "tiny"

func HandleTextMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	text := strings.ToLower(message.Text)

	logFilePath := "storage/logs/text_handler.log"

	if _, err := os.Stat("storage/logs"); os.IsNotExist(err) {
		err := os.MkdirAll("storage/logs", 0755)
		if err != nil {
			fmt.Println("Failed to create log directory:", err)
			return
		}
	}

	logInstance, err := logger.GetLogger(logFilePath)
	if err != nil {
		fmt.Println("Error initializing logger:", err)
		return
	}
	defer logInstance.Close()

	logInstance.Info(fmt.Sprintf("Received message: %s", text))

	if strings.Contains(text, "model: ") {
		parts := strings.Split(text, ":")
		if len(parts) == 2 {
			model := strings.TrimSpace(parts[1])

			if model == "tiny" || model == "large" {
				SelectedModel = model

				response := fmt.Sprintf("Model successfully changed to: %s", SelectedModel)
				msg := tgbotapi.NewMessage(message.Chat.ID, response)
				bot.Send(msg)

				logInstance.Info(fmt.Sprintf("Model changed to: %s", SelectedModel))
				return
			}
		}
	}

	logInstance.Error("Invalid model received or format incorrect.")
}
