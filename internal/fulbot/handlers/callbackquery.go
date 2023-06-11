package handlers

import (
	"fulbot/internal/gateways/telegram"

	telegramapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackQueryHandler struct {
	bot *telegram.Bot
}

func NewCallbackQuery(bot *telegram.Bot) *CallbackQueryHandler {
	return &CallbackQueryHandler{
		bot: bot,
	}
}

func (c *CallbackQueryHandler) Handle(update telegramapi.Update) error {
	// Respond to the callback query, telling Telegram to show the user
	// a message with the data received.

	callback := telegramapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := c.bot.Client.Request(callback); err != nil {
		panic(err)
	}

	// And finally, send a message containing the data received.
	msg := telegramapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
	if _, err := c.bot.Client.Send(msg); err != nil {
		panic(err)
	}

	return nil
}
