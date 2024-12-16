package tgbot_test

import (
	"testing"

	"github.com/alaa-aqeel/bot-orders-demo/src/pkg/tgbot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TestTgBotAddCommands(t *testing.T) {
	commands := tgbot.NewCommands()
	commands.Adds(
		tgbot.Command{
			Command:     "start",
			Description: "Starts the bot",
			Callback:    func(bot *tgbotapi.BotAPI, update tgbotapi.Update) {},
		},
		tgbot.Command{
			Command:     "say-hello",
			Description: "Say hello world",
			Callback:    func(bot *tgbotapi.BotAPI, update tgbotapi.Update) {},
		},
	)
	expectedCommands := []tgbotapi.BotCommand{
		{Command: "start", Description: "Starts the bot"},
		{Command: "say-hello", Description: "Say hello world"},
	}
	registeredCommands := commands.GetBotCommands()
	if len(registeredCommands) != len(expectedCommands) {
		t.Fatalf("Expected %d commands, got %d", len(expectedCommands), len(registeredCommands))
	}

	for i, cmd := range registeredCommands {
		if cmd.Command != expectedCommands[i].Command {
			t.Errorf("Command mismatch: expected %v, got %v", expectedCommands[i], cmd)
		}
		// check callback command
		if commands.GetCommandCallback(cmd.Command) == nil {
			t.Errorf("Command mismatch (%v): callback nill", cmd.Command)
		}
	}
}

func TestTgBotCommandsCallback(t *testing.T) {
	mockBot := &tgbotapi.BotAPI{}
	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "/helloword",
			Chat: &tgbotapi.Chat{ID: 12345},
		},
	}
	commands := tgbot.NewCommands(
		tgbot.Command{
			Command:     "hello",
			Description: "Say hello world",
			Callback: func(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

			},
		},
	)
	commandCallback := commands.GetCommandCallback("hello")
	if commandCallback == nil {
		t.Errorf("Command mismatch (hello): callback nill")
	}
	commandCallback(mockBot, update)
}
