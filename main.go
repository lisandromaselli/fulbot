package main

import (
	"fulbot/internal/fulbot"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	app, err := fulbot.NewApp()
	if err != nil {
		log.Error().Msgf("Error creating the app, shutting down: %s", err)
		return
	}

	err = app.Run()
	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}
}
