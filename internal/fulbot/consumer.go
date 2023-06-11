package fulbot

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type UpdateManger interface {
	ProcessUpdate(telegram.Update)
}

type UpdateConsumer struct {
	manager        UpdateManger
	UpdatesChannel telegram.UpdatesChannel
}

func NewUpdateConsumer(
	channel telegram.UpdatesChannel,
	manager UpdateManger,
) *UpdateConsumer {
	return &UpdateConsumer{
		UpdatesChannel: channel,
		manager:        manager,
	}
}

func (c UpdateConsumer) Run() {
	log.Info().Msg("Starting event listening")

	for update := range c.UpdatesChannel {
		c.manager.ProcessUpdate(update)
	}
}
