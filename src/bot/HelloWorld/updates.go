package helloworld

import (
	"github.com/alaa-aqeel/bot-orders-demo/src/pkg/tgbot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func botUpdates(bot *tgbotapi.BotAPI, commands *tgbot.BotCommands) {
	tgbot.HandleUpdates(bot, func(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
		if update.Message.IsCommand() {
			commands.HandleCommand(bot, update)
		}
	})
}
