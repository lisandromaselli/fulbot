package fulbot

import (
	"errors"
	"fulbot/internal/fulbot/handlers/callbacksquery"
	"fulbot/internal/fulbot/handlers/commands"
	"fulbot/internal/pkg/telegram"

	"github.com/rs/zerolog/log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

	bot, err := telegram.NewTelegramBot(config.TelegramToken)
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

func openUpdatesConnection(bot *telegram.TelegramBot, config *Config) (updates tgbotapi.UpdatesChannel, err error) {
	if config.ConnectionStrategy == "WEBHOOK" {
		log.Info().Msg("Starting telegram webhook")
		updates, err = bot.StartTelegramWebHook(config.TelegramWebHookConfig())
		return
	}

	if config.ConnectionStrategy == "POOLING" {
		log.Info().Msg("Starting telegram pooling")
		updates, err = bot.StartTegramDaemon()
		return
	}

	err = errors.New("Invalid connection strategy")
	return
}
