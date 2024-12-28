// bot/handlers/command_handler.go
package handlers

import (
	"fmt"

	"tg-whisper-bot/bot/utils/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	ModelTiny  = "tiny"
	ModelBase  = "base"
	ModelSmall = "small"
	ModelMedium = "medium"
	ModelLarge = "large"
)

var SelectedModel = ModelTiny

func HandleModelCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, logInstance *logger.Logger) {
	buttons := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("Tiny", ModelTiny),
		tgbotapi.NewInlineKeyboardButtonData("Base", ModelBase),
		tgbotapi.NewInlineKeyboardButtonData("Small", ModelSmall),
		tgbotapi.NewInlineKeyboardButtonData("Medium", ModelMedium),
		tgbotapi.NewInlineKeyboardButtonData("Large", ModelLarge),
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Choose the model:")
	msg.ReplyMarkup = keyboard

	if _, err := bot.Send(msg); err != nil {
		logInstance.Error(fmt.Sprintf("Error sending model selection message: %v", err))
		return
	}

	logInstance.Info("Model selection keyboard sent.")
}

func HandleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery, logInstance *logger.Logger) {
	model := callback.Data

	if model == ModelTiny || model == ModelBase || model == ModelSmall || model == ModelMedium || model == ModelLarge {
		SelectedModel = model

		response := fmt.Sprintf("Model successfully changed to: %s", SelectedModel)
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, response)
		if _, err := bot.Send(msg); err != nil {
			logInstance.Error(fmt.Sprintf("Error sending message: %v", err))
		}

		callbackResponse := tgbotapi.NewCallback(callback.ID, fmt.Sprintf("Selected: %s", model))
		if _, err := bot.Request(callbackResponse); err != nil {
			logInstance.Error(fmt.Sprintf("Error responding to callback: %v", err))
		}

		logInstance.Info(fmt.Sprintf("Model changed to: %s", SelectedModel))
		return
	}

	errorResponse := "Invalid model selected. Please choose a valid model."
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, errorResponse)
	if _, err := bot.Send(msg); err != nil {
		logInstance.Error(fmt.Sprintf("Error sending invalid model response: %v", err))
	}

	logInstance.Error(fmt.Sprintf("Invalid model selected: %s", model))
}

func HandleInfoCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, logInstance *logger.Logger) {
	logInstance.Info(fmt.Sprintf("User %s (ID: %d) issued the /info command", message.From.UserName, message.From.ID))

	botInfo := fmt.Sprintf(
		`🤖 *Voice-to-Text Bot Information* 🤖
		Hello, *%s*! 
		I am a bot designed to convert voice messages into text. Here are some of my features:
		- 🎙️ Convert voice messages to text using advanced speech-to-text models.
		- 🌐 Support for Russian language.
		- 📜 Easy command-based usage.
		- 🔧 Configurable to fit your needs.
		
		💡 *Usage:*
		1️⃣ Send me a voice message, and I'll transcribe it into text.
		2️⃣ Use commands like /info, and /model for guidance.
		
		⚙️ Currently integrated with *Whisper API* for transcription.
		
		If you encounter issues or have suggestions, please visit the [GitHub Repository](https://github.com/AlmirSai/MVP_bot_voice_to_text_dl) or contact the developer.
		
		Thank you for using this bot!`,
				message.From.FirstName,
	)

	btnGitHub := tgbotapi.NewInlineKeyboardButtonURL("GitHub Repository", "https://github.com/AlmirSai/MVP_bot_voice_to_text_dl")

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(btnGitHub),
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, botInfo)
	msg.ParseMode = "Markdown"
	msg.DisableWebPagePreview = true
	msg.ReplyMarkup = keyboard

	if _, err := bot.Send(msg); err != nil {
		logInstance.Error(fmt.Sprintf("Failed to send /info response: %v", err))
	}
}
