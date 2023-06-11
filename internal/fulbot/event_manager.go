package fulbot

import (
	"errors"
	"fulbot/internal/fulbot/handlers"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var errDuplicateCommandHandler = errors.New("duplicated command handler")

type TelegramEventManager struct {
	commandMap       map[string]handlers.TelegramCommandHandler
	callbackQHandler handlers.TelegramUpdateHandler
}

func NewTelegramEventManager(cbh handlers.TelegramUpdateHandler) *TelegramEventManager {
	return &TelegramEventManager{
		commandMap:       map[string]handlers.TelegramCommandHandler{},
		callbackQHandler: cbh,
	}
}

func (em *TelegramEventManager) AddCommandHandler(handler handlers.TelegramCommandHandler) error {
	_, ok := em.commandMap[handler.Suscribe()]
	if ok {
		return errDuplicateCommandHandler
	}

	em.commandMap[handler.Suscribe()] = handler

	return nil
}

func (em *TelegramEventManager) ProcessUpdate(update telegram.Update) {
	commands := []struct {
		check   func() bool
		command func() string
	}{
		{
			check:   func() bool { return update.Message != nil && update.Message.IsCommand() },
			command: update.Message.Command},
		{check: func() bool { return update.EditedMessage != nil && update.EditedMessage.IsCommand() },
			command: update.EditedMessage.Command},
		{check: func() bool { return update.ChannelPost != nil && update.ChannelPost.IsCommand() },
			command: update.ChannelPost.Command},
		{check: func() bool { return update.EditedChannelPost != nil && update.EditedChannelPost.IsCommand() },
			command: update.EditedChannelPost.Command},
	}

	for _, cmd := range commands {
		if cmd.check() {
			if h, ok := em.commandMap[cmd.command()]; ok {
				// TODO: handle errors of update handlers
				_ = h.Handle(update)
				return
			}
		}
	}

	if update.CallbackQuery != nil {
		_ = em.callbackQHandler.Handle(update)
		return
	}
}
