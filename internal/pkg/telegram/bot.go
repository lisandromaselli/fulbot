package telegram

import (
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &TelegramBot{
		bot: bot,
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

	return t.bot.GetUpdatesChan(u), nil

}
func (t *TelegramBot) StartTelegramWebHook(port, webhookUrl string) (tgbotapi.UpdatesChannel, error) {
	wh, err := tgbotapi.NewWebhook(webhookUrl)
	if err != nil {
		return nil, err
	}

	err = suscribeWebhook(t.bot, wh)
	if err != nil {
		return nil, err
	}

	updates := t.bot.ListenForWebhook("/webhooks")
	go http.ListenAndServe("0.0.0.0:"+port, nil)

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

func (t *TelegramBot) Send(msg tgbotapi.MessageConfig) {
	t.bot.Send(msg)
}
