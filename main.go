package main

import (
	"fulbot/internal/fulbot"

	"log"
)

func main() {
	app, err := fulbot.NewApp()
	if err != nil {
		log.Printf("Error creating the app, shutting down")
		return
	}
	app.Run()
}
