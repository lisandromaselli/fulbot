package callbacksquery

import (
	"fulbot/internal/fulbot/handlers"
	"fulbot/internal/gateways/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

func NewCallbackQueryMatch(bot *telegram.Bot) handlers.UpdateHandler {
	return handlers.NewCallbackQueryHandler(dummyPattern, func(update tgbotapi.Update) error {
		// Respond to the callback query, telling Telegram to show the user
		// a message with the data received.
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
		if _, err := bot.Client.Request(callback); err != nil {
			panic(err)
		}

		// And finally, send a message containing the data received.
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
		if _, err := bot.Client.Send(msg); err != nil {
			panic(err)
		}

		return nil
	})
}

func dummyPattern(update tgbotapi.Update) bool {
	data := update.CallbackData()
	log.Info().Msgf("data: %v\n", data)

	return true
}
