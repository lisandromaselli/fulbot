package telegram

import (
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Client *tgbotapi.BotAPI
}

type WebHookConfig struct {
	Port, Domain, WebhookSecretPath string
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	bot.Debug = true
	if err != nil {
		return nil, err
	}
	return &TelegramBot{
		Client: bot,
	}, nil
}

func (t *TelegramBot) StartTegramDaemon() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	u.AllowedUpdates = []string{
		tgbotapi.UpdateTypeMessage,
		tgbotapi.UpdateTypeEditedMessage,
		tgbotapi.UpdateTypeChannelPost,
		tgbotapi.UpdateTypeEditedChannelPost,
		tgbotapi.UpdateTypeInlineQuery,
		tgbotapi.UpdateTypeChosenInlineResult,
		tgbotapi.UpdateTypeCallbackQuery,
		tgbotapi.UpdateTypeShippingQuery,
		tgbotapi.UpdateTypePreCheckoutQuery,
		tgbotapi.UpdateTypePoll,
		tgbotapi.UpdateTypePollAnswer,
		tgbotapi.UpdateTypeMyChatMember,
		tgbotapi.UpdateTypeChatMember,
	}

	return t.Client.GetUpdatesChan(u), nil
}

func (t *TelegramBot) StartTelegramWebHook(config WebHookConfig) (tgbotapi.UpdatesChannel, error) {
	wh, err := tgbotapi.NewWebhook(config.Domain + config.WebhookSecretPath)
	if err != nil {
		return nil, err
	}

	err = suscribeWebhook(t.Client, wh)
	if err != nil {
		return nil, err
	}
	log.Print("Successfully suscribed to the webhook")

	updates := t.Client.ListenForWebhook(config.WebhookSecretPath)
	go http.ListenAndServe(":"+config.Port, nil)

	return updates, nil
}

func suscribeWebhook(bot *tgbotapi.BotAPI, wh tgbotapi.WebhookConfig) error {
	_, err := bot.Request(wh)
	if err != nil {
		return err
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		return err
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	return nil
}
