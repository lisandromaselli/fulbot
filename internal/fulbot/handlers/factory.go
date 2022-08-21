package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type CommandHandler struct {
	command      string
	callbackFunc func(tgbotapi.Update) error
}

func NewCommandHandler(command string, callbackFunc func(tgbotapi.Update) error) CommandHandler {
	return CommandHandler{
		command:      command,
		callbackFunc: callbackFunc,
	}
}

func (c CommandHandler) CheckUpdate(update tgbotapi.Update) bool {
	if update.Message.IsCommand() && update.Message.Command() == c.command {
		return true
	}
	return false
}

func (c CommandHandler) HandleUpdate(update tgbotapi.Update) error {
	return c.callbackFunc(update)
}
