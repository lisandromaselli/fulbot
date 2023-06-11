package handlers

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramUpdateHandler interface {
	Handle(telegram.Update) error
}

type TelegramCommandHandler interface {
	TelegramUpdateHandler
	Suscribe() string
}
