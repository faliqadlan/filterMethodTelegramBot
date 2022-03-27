package main

import (
	"fm/config"

	"github.com/labstack/gommon/log"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	var config = config.GetConfig()

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		panic(err)
	}

	updatedConfig := tgbotapi.NewUpdate(0)

	log.Info(updatedConfig)

	updatedConfig.Timeout = 30
	log.Info(updatedConfig)

	updates := bot.GetUpdatesChan(updatedConfig)
	log.Info(updates)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		log.Info(msg)

		msg.ReplyToMessageID = update.Message.MessageID

		log.Info(msg)

		if _, err := bot.Send(msg); err != nil {

			panic(err)
		}

	}

	bot.Debug = true
}
