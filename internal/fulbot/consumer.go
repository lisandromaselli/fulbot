package fulbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UpdateConsumer struct {
	manager        EventManager
	UpdatesChannel tgbotapi.UpdatesChannel
}

func NewUpdateConsumer(channel tgbotapi.UpdatesChannel, manager EventManager) *UpdateConsumer {
	return &UpdateConsumer{
		UpdatesChannel: channel,
		manager:        manager,
	}
}

func (c UpdateConsumer) Run() {
	for update := range c.UpdatesChannel {
		c.manager.ProcessUpdate(update)
	}
}
