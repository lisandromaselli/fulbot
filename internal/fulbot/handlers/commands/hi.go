package commands

import (
	"fulbot/internal/fulbot/handlers"
	"fulbot/internal/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewHiCommand(bot *telegram.Bot) handlers.UpdateHandler {
	return handlers.NewCommandHandler("hi", func(u tgbotapi.Update) error {
		msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Hi")
		msg.ReplyToMessageID = u.Message.MessageID

		bot.Client.Send(msg)
		return nil
	})
}
