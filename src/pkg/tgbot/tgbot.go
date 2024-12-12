package tgbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotConfig struct {
	AccessToken string `json:"access_token"`
}

func NewBot(config BotConfig) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(config.AccessToken)
	if err != nil {
		log.Panic("[BOT-ERROR]: ", err)
	}
	log.Printf("[BOT]Authorized on account %s", bot.Self.UserName)

	return bot
}
