package main

import (
	"github.com/alaa-aqeel/bot-orders-demo/src/bot"
	"github.com/alaa-aqeel/bot-orders-demo/src/helpers"
)

func main() {
	helpers.LoadEnvFile()

	bot.NewBot()
}
