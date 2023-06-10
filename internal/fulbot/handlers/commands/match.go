package commands

import (
	"fulbot/internal/fulbot/handlers"
	"fulbot/internal/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("create", "match_create"),
		tgbotapi.NewInlineKeyboardButtonData("cancel", "match_cancel"),
		tgbotapi.NewInlineKeyboardButtonData("update", "match_update"),
	),
)

func NewMatchCommand(bot *telegram.Bot) handlers.UpdateHandler {
	return handlers.NewCommandHandler("match", func(u tgbotapi.Update) error {
		msg := tgbotapi.NewMessage(u.Message.Chat.ID, u.Message.Text)
		msg.ReplyMarkup = numericKeyboard
		if _, err := bot.Client.Send(msg); err != nil {
			return err
		}
		return nil
	})
}
