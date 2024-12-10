package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg-whisper-bot/bot/config"
	"tg-whisper-bot/bot/handlers"
)

func main() {
	if err := config.CheckDependencies(); err != nil {
		log.Fatalf("Missing dependencies: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(config.GetToken())
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

		if update.Message.Text != "" {
			handlers.HandleTextMessage(bot, update.Message)
			continue
		}

		if update.Message.Voice != nil {
			handlers.HandleVoiceMessage(bot, update.Message)
		}
	}
}