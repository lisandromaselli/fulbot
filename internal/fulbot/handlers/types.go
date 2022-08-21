package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UpdateHandler interface {
	CheckUpdate(tgbotapi.Update) bool
	HandleUpdate(tgbotapi.Update) error
}
