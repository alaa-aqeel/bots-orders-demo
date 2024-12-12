package tgbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HandleUpdate func(bot *tgbotapi.BotAPI, update tgbotapi.Update)

func HandleUpdates(bot *tgbotapi.BotAPI, handleUpdate HandleUpdate) {
	updates := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:  0,
		Limit:   0,
		Timeout: 60,
	})
	for update := range updates {
		if update.Message == nil {
			continue
		}
		handleUpdate(bot, update)
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}
