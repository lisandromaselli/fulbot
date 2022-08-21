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
	if update.Message != nil &&
		update.Message.IsCommand() &&
		update.Message.Command() == c.command {
		return true
	}
	return false
}

func (c CommandHandler) HandleUpdate(update tgbotapi.Update) error {
	return c.callbackFunc(update)
}

type CallbackQueryHandler struct {
	patternFunc  func(tgbotapi.Update) bool
	callbackFunc func(tgbotapi.Update) error
}

func NewCallbackQueryHandler(patternFunc func(tgbotapi.Update) bool,
	callbackFunc func(tgbotapi.Update) error) CallbackQueryHandler {
	return CallbackQueryHandler{
		patternFunc:  patternFunc,
		callbackFunc: callbackFunc,
	}
}

func (c CallbackQueryHandler) CheckUpdate(update tgbotapi.Update) bool {
	if update.CallbackQuery != nil && c.patternFunc(update) {
		return true
	}
	return false
}

func (c CallbackQueryHandler) HandleUpdate(update tgbotapi.Update) error {
	return c.callbackFunc(update)
}
