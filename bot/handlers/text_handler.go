package handlers

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var SelectedModel = "tiny"

func HandleTextMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	text := strings.ToLower(message.Text)

	if strings.Contains(text, "model: ") {
		parts := strings.Split(text, ":")
		if len(parts) == 2 {
			model := strings.TrimSpace(parts[1])

			if model == "tiny" || model == "large" {
				SelectedModel = model

				response := fmt.Sprintf("Model successfully changed to: %s", SelectedModel)
				msg := tgbotapi.NewMessage(message.Chat.ID, response)
				bot.Send(msg)
				return
			}
		}
	}
}