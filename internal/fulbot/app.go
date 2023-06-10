package fulbot

import (
	"errors"
	"fmt"

	"fulbot/internal/fulbot/handlers/callbacksquery"
	"fulbot/internal/fulbot/handlers/commands"
	"fulbot/internal/gateways/telegram"

	"github.com/rs/zerolog/log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var errInvalidStrategy = errors.New("invalid connection strategy")

type App struct{}

func NewApp() (App, error) {
	return App{}, nil
}

func (app *App) Run() {
	config, err := LoadConfig()
	if err != nil {
		log.Error().Err(err).Msg("Error loading config")
		return
	}

	bot, err := telegram.NewBot(config.TelegramToken)
	if err != nil {
		log.Error().Err(err).Msg("Error building telegram bot")
		return
	}

	updates, err := openUpdatesConnection(bot, config)
	if err != nil {
		log.Error().Err(err).Msg("Error opening updates connection")
		return
	}

	manager := NewUpdateManager()
	manager.AddHandler(commands.NewHiCommand(bot))
	manager.AddHandler(commands.NewMatchCommand(bot))
	manager.AddHandler(callbacksquery.NewCallbackQueryMatch(bot))

	log.Info().Msg("Starting event listening")

	consumer := NewUpdateConsumer(updates, manager)
	consumer.Run()
}

func openUpdatesConnection(bot *telegram.Bot, config *Config) (tgbotapi.UpdatesChannel, error) {
	if config.ConnectionStrategy == "WEBHOOK" {
		u, err := bot.StartTelegramWebHook(config.TelegramWebHookConfig())
		if err != nil {
			return nil, fmt.Errorf("cannot update webhook connection %w", err)
		}

		return u, nil
	}

	if config.ConnectionStrategy == "POOLING" {
		u, err := bot.StartTegramDaemon()
		if err != nil {
			return nil, fmt.Errorf("cannot update pooling connection %w", err)
		}

		return u, nil
	}

	return nil, errInvalidStrategy
}
