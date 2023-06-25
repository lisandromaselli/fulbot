package main

import (
	"fulbot/internal/fulbot"
	"fulbot/internal/gateways/alquilatucancha"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	id, err := alquilatucancha.ExtractBuildID()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	params := alquilatucancha.GetSportClubParams{
		Day:     "2023-06-17",
		BuildID: id,
	}

	resp, errA := alquilatucancha.GetSportClubAvailability(params)
	if errA != nil {
		log.Error().Msgf("Error creating the app, shutting down: %s", errA)
		return
	}

	if bookings := resp.PageProps.BookingsBySport.Num2; bookings != nil {
		for _, b := range bookings {
			for _, c := range b.Courts {
				for _, a := range c.Available {
					log.Info().Msgf("%s %s %v", b.Name, c.Name, a)
				}
			}
		}
	}

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
