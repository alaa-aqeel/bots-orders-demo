package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandCallback func(bot *tgbotapi.BotAPI, update tgbotapi.Update)

type Command struct {
	Command     string
	Description string
	Callback    CommandCallback
}

type BotCommand struct {
	botCommands []tgbotapi.BotCommand
	commands    map[string]CommandCallback
}

func NewCommands(cmd ...Command) BotCommand {
	bot := &BotCommand{
		botCommands: []tgbotapi.BotCommand{},
		commands:    map[string]CommandCallback{},
	}

	return bot.Adds(cmd...)
}

func (c BotCommand) Adds(cmd ...Command) BotCommand {
	if len(cmd) <= 0 {
		return c
	}
	for _, cmd := range cmd {
		c.botCommands = append(c.botCommands, tgbotapi.BotCommand{
			Command:     cmd.Command,
			Description: cmd.Description,
		})
		c.commands[cmd.Command] = cmd.Callback
	}

	return c
}

func (c BotCommand) NewSetMyCommands() tgbotapi.SetMyCommandsConfig {

	return tgbotapi.NewSetMyCommands(c.botCommands...)
}

func (c BotCommand) GetCommandCallback(cmd string) CommandCallback {
	callback, ok := c.commands[cmd]
	if ok {
		return callback
	}

	return nil
}

func (c BotCommand) HandleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	callback := c.GetCommandCallback(update.Message.Command())
	if callback != nil {
		callback(bot, update)
	}
}
