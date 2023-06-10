package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/rs/zerolog/log"
)

type Bot struct {
	Client *tgbotapi.BotAPI
}

type WebHookConfig struct {
	Port, Domain, WebhookSecretPath string
}

func NewBot(token string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("cannto build bot: %w", err)
	}

	bot.Debug = true

	return &Bot{
		Client: bot,
	}, nil
}

func (t *Bot) StartTegramDaemon() (tgbotapi.UpdatesChannel, error) {
	log.Info().Msg("Starting telegram pooling")

	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60
	update.AllowedUpdates = []string{
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

	return t.Client.GetUpdatesChan(update), nil
}

func (t *Bot) StartTelegramWebHook(config WebHookConfig) (tgbotapi.UpdatesChannel, error) {
	log.Info().Msg("Starting telegram webhook")

	wh, err := tgbotapi.NewWebhook(config.Domain + config.WebhookSecretPath)
	if err != nil {
		return nil, fmt.Errorf("cannot start webhook server %w", err)
	}

	err = suscribeWebhook(t.Client, wh)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("Successfully suscribed to the webhook")

	updates := t.Client.ListenForWebhook(config.WebhookSecretPath)

	return updates, nil
}

func suscribeWebhook(bot *tgbotapi.BotAPI, wh tgbotapi.WebhookConfig) error {
	_, err := bot.Request(wh)
	if err != nil {
		return fmt.Errorf("cannto suscribe webhook: %w", err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		return fmt.Errorf("cannto get webhook info: %w", err)
	}

	if info.LastErrorDate != 0 {
		log.Info().Msgf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	return nil
}
