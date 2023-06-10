package main

import (
	"fulbot/internal/fulbot"

	"github.com/rs/zerolog/log"
)

func main() {
	app, err := fulbot.NewApp()
	if err != nil {
		log.Error().Msg("Error creating the app, shutting down")
		return
	}

	app.Run()
}
