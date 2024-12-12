package main

import (
	helloworld "github.com/alaa-aqeel/bot-orders-demo/src/bot/HelloWorld"
	"github.com/alaa-aqeel/bot-orders-demo/src/helpers"
)

func main() {
	helpers.LoadEnvFile()

	helloworld.NewBot()
}
