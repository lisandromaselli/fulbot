package fulbot

import (
	"errors"
	"fmt"

	"fulbot/internal/fulbot/handlers"
	"fulbot/internal/fulbot/handlers/commands"
	"fulbot/internal/gateways/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var errInvalidStrategy = errors.New("invalid connection strategy")

type App struct{}

func NewApp() (*App, error) {
	return &App{}, nil
}

func (app *App) Run() error {
	config, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading config %w", err)
	}

	bot, err := telegram.NewBot(config.TelegramToken)
	if err != nil {
		return fmt.Errorf("error building telegram bot %w", err)
	}

	updates, err := openUpdatesConnection(bot, config)
	if err != nil {
		return fmt.Errorf("error opening updates connection %w", err)
	}

	manager := NewTelegramEventManager(handlers.NewCallbackQuery(bot))

	err = manager.AddCommandHandler(commands.NewHiCommand(bot))
	if err != nil {
		return err
	}

	err = manager.AddCommandHandler(commands.NewMatchCommand(bot))
	if err != nil {
		return err
	}

	consumer := NewUpdateConsumer(updates, manager)
	consumer.Run()

	return nil
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
