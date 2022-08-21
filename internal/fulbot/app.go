package fulbot

import (
	"fulbot/internal/fulbot/handlers"
	"fulbot/internal/pkg/telegram"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct{}

func NewApp() (App, error) {
	return App{}, nil
}

func (app *App) Run() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("Telegram token not found")
	}
	webhookUrl := os.Getenv("WEBHOOK_FULL_URL")
	port := os.Getenv("PORT")

	bot, err := telegram.NewTelegramBot(token)
	if err != nil {
		return
	}

	updates, err := bot.StartTelegramWebHook(webhookUrl, port)
	if err != nil {
		return
	}

	manager := NewUpdateManager()
	manager.AddHandler(handlers.NewCommandHandler("hi", func(u tgbotapi.Update) error {
		log.Printf("[%s] %s", u.Message.From.UserName, u.Message.Text)

		msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Ariel trolo")
		msg.ReplyToMessageID = u.Message.MessageID

		bot.Send(msg)
		return nil
	}))

	consumer := NewUpdateConsumer(updates, manager)
	consumer.Run()
}
