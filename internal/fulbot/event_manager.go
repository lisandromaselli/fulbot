package fulbot

import (
	"fulbot/internal/fulbot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type EventManager interface {
	AddHandler(handler handlers.UpdateHandler)
	ProcessUpdate(update tgbotapi.Update)
}

type UpdateManager struct {
	handlerList []handlers.UpdateHandler
}

func NewUpdateManager() *UpdateManager {
	return &UpdateManager{
		handlerList: []handlers.UpdateHandler{},
	}
}

func (m *UpdateManager) AddHandler(handler handlers.UpdateHandler) {
	m.handlerList = append(m.handlerList, handler)
}

func (m *UpdateManager) ProcessUpdate(update tgbotapi.Update) {
	for _, h := range m.handlerList {
		if h.CheckUpdate(update) {
			_ = h.HandleUpdate(update)
		}
	}
}
