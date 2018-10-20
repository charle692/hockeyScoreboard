package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Domain - NHL API
const Domain = "https://statsapi.web.nhl.com"

// Game - Contains full game data
type Game struct {
	Link     string   `json:"link"`
	GameData GameData `json:"gameData"`
	LiveData LiveData `json:"liveData"`
}

// RetrieveGame - Retrieves game information given the game link
func RetrieveGame(link string) (*Game, error) {
	game := &Game{}
	resp, err := http.Get(Domain + link)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, game); err != nil {
		return nil, err
	}

	return game, nil
}

// ListenForGameUpdates - iteratively retrieve game information
func (g *Game) ListenForGameUpdates(newTeamSelected chan bool, gameStarted *bool) {
	ticker := time.NewTicker(time.Second * 2)

	for range ticker.C {
		select {
		case <-newTeamSelected:
			*gameStarted = false
			ticker.Stop()
			return
		default:
			g, err := RetrieveGame(g.Link)

			if err != nil {
				log.Println("An error occurred while retrieving game data")
			}

			log.Printf("Home Goals: %d\n", g.LiveData.Linescore.Teams.Home.Goals)
			log.Printf("Away Goals: %d\n", g.LiveData.Linescore.Teams.Away.Goals)
			log.Printf("Period: %d\n", g.LiveData.Linescore.CurrentPeriod)
			log.Printf("Time Remaining: %s\n", g.LiveData.Linescore.CurrentPeriodTimeRemaining)
			log.Printf("Home Shots on Goal: %d\n", g.LiveData.Linescore.Teams.Home.ShotsOnGoal)
			log.Printf("Away Shots on Goal: %d\n", g.LiveData.Linescore.Teams.Away.ShotsOnGoal)
		}
	}
}
