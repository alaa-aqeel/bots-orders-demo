package helloworld

import (
	"github.com/alaa-aqeel/bot-orders-demo/src/pkg/tgbot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func botCommands(bot *tgbotapi.BotAPI) *tgbot.BotCommands {
	comands := tgbot.NewCommands(
		tgbot.Command{
			Command:     "start",
			Description: "Starts the bot",
			Callback: func(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

			},
		},
		tgbot.Command{
			Command:     "helloword",
			Description: "Say hello world",
			Callback: func(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Hello World"))
			},
		},
	)
	bot.Send(comands.NewSetMyCommands())

	return comands
}
