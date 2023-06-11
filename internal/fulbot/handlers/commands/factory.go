package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type CommandHandler struct {
	command      string
	callbackFunc func(tgbotapi.Update) error
}

func NewCommandHandler(
	command string,
	callbackFunc func(tgbotapi.Update) error,
) CommandHandler {
	return CommandHandler{
		command:      command,
		callbackFunc: callbackFunc,
	}
}
func (c CommandHandler) Suscribe() string {
	return c.command
}

func (c CommandHandler) Handle(update tgbotapi.Update) error {
	return c.callbackFunc(update)
}
