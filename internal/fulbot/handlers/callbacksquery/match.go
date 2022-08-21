package callbacksquery

import (
	"fmt"
	"fulbot/internal/fulbot/handlers"
	"fulbot/internal/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewCallbackQueryMatch(bot *telegram.TelegramBot) handlers.UpdateHandler {
	return handlers.NewCallbackQueryHandler(dummyPattern, func(u tgbotapi.Update) error {
		// Respond to the callback query, telling Telegram to show the user
		// a message with the data received.
		callback := tgbotapi.NewCallback(u.CallbackQuery.ID, u.CallbackQuery.Data)
		if _, err := bot.Client.Request(callback); err != nil {
			panic(err)
		}

		// And finally, send a message containing the data received.
		msg := tgbotapi.NewMessage(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Data)
		if _, err := bot.Client.Send(msg); err != nil {
			panic(err)
		}
		return nil
	})
}

func dummyPattern(update tgbotapi.Update) bool {
	data := update.CallbackData()
	fmt.Printf("data: %v\n", data)
	return true
}
