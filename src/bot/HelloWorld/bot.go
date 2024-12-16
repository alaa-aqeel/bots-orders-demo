package helloworld

import (
	"os"

	"github.com/alaa-aqeel/bot-orders-demo/src/pkg/tgbot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBot() *tgbotapi.BotAPI {
	bot := tgbot.NewBot(tgbot.BotConfig{
		AccessToken: os.Getenv("BOT_ACCESS_TOKEN"),
	})
	if bot != nil {
		cmd := botCommands(bot)
		botUpdates(bot, cmd)
	}
	return bot
}
