// bot/cmd/main.go
package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg-whisper-bot/bot/config"
	"tg-whisper-bot/bot/handlers"
	"tg-whisper-bot/bot/utils/logger"
)

// main initializes the bot, sets up the logger, and starts the main loop.
//
// It reads the Telegram bot token from the .env file, checks for dependencies,
// and initializes the Telegram bot API.
//
// It then sets up the logger and starts the main loop, which listens for
// incoming updates, handles callbacks (model changes), and processes text and
// voice messages.
func main() {
	token, err := config.GetToken(".env")
	if err != nil {
		log.Fatalf("Error reading token: %v", err)
	}

	if err := config.CheckDependencies(); err != nil {
		log.Fatalf("Missing dependencies: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Error initializing bot: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	logFilePath := "storage/logs/bot.log"
	logInstance, err := logger.GetLogger(logFilePath)
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer logInstance.Close()

	for update := range updates {
		if update.CallbackQuery != nil {
			handlers.HandleCallback(bot, update.CallbackQuery, logInstance)
			continue
		}

		if update.Message == nil {
			continue
		}

		if update.Message.Text != "" {
			handlers.HandleTextMessage(bot, update.Message)
			continue
		}

		if update.Message.Voice != nil {
			handlers.HandleVoiceMessage(bot, update.Message)
		}
	}
}
