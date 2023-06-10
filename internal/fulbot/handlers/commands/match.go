package commands

import (
	"fmt"

	"fulbot/internal/fulbot/handlers"
	"fulbot/internal/gateways/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewMatchCommand(bot *telegram.Bot) handlers.UpdateHandler {
	numericKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("create", "match_create"),
			tgbotapi.NewInlineKeyboardButtonData("cancel", "match_cancel"),
			tgbotapi.NewInlineKeyboardButtonData("update", "match_update"),
		),
	)

	return handlers.NewCommandHandler("match", func(u tgbotapi.Update) error {
		msg := tgbotapi.NewMessage(u.Message.Chat.ID, u.Message.Text)
		msg.ReplyMarkup = numericKeyboard
		if _, err := bot.Client.Send(msg); err != nil {
			return fmt.Errorf("failed to publish telegram message: %w", err)
		}

		return nil
	})
}
