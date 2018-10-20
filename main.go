package main

import (
	"log"

	"github.com/charle692/hockeyScoreBoard/models"
)

// Domain - NHP API Domain
const Domain = "https://statsapi.web.nhl.com"

func main() {
	// database.Conn.AutoMigrate(&models.Message{}, &models.MediaItem{}, &models.Contact{}, &models.Conversation{})
	gameStarted := false
	newTeamSelected := make(chan bool)

	gameLink := waitForGameToStart(gameStarted)

	select {
	case link := <-gameLink:
		if !gameStarted {
			gameStarted = true
			game, err := models.RetrieveGame(link)

			if err != nil {
				log.Println("An error occurred while retrieving initial game data")
			} else {
				go game.ListenForGameUpdates(newTeamSelected, &gameStarted)
			}
		}
	}
}
