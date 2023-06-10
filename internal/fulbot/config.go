package fulbot

import (
	"fmt"

	"fulbot/internal/gateways/telegram"

	"github.com/openware/pkg/ika"
)

type Config struct {
	TelegramToken      string `env:"TELEGRAM_TOKEN" env-required:"true"`
	Port               string `env:"PORT"`
	Domain             string `env:"DOMAIN"`
	WebhookSecretPath  string `env:"WEBHOOOK_SECRET_PATH"`
	ConnectionStrategy string `env:"CONNECTION_STRATEGY" env-default:"POOLING"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		TelegramToken:      "",
		Port:               "",
		Domain:             "",
		WebhookSecretPath:  "",
		ConnectionStrategy: "",
	}
	if err := ika.ReadConfig("", cfg); err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	return cfg, nil
}

func (config Config) TelegramWebHookConfig() telegram.WebHookConfig {
	return telegram.WebHookConfig{
		Port:              config.Port,
		Domain:            config.Domain,
		WebhookSecretPath: config.WebhookSecretPath,
	}
}
