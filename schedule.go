package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/charle692/hockeyScoreBoard/models"
)

// SchedulePath - Path to schedule
const SchedulePath = "/api/v1/schedule"

func waitForGameToStart(gameStarted bool) <-chan string {
	gameLinkChan := make(chan string)
	go func() {
		for range time.NewTicker(time.Second * 30).C {
			if !gameStarted {
				isTeamPlayingToday(gameLinkChan)
			}
		}
	}()
	return gameLinkChan
}

func isTeamPlayingToday(gameLinkChan chan string) {
	gameDate := &models.GameDate{}
	resp, err := http.Get(Domain + SchedulePath + "?teamId=" + "23")

	if err != nil {
		fmt.Printf("An error occured while retrieving today's game data")
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("An error occured while retrieving today's game data")
	}

	json.Unmarshal(body, gameDate)

	if gameDate.TotalGames > 0 {
		gameLinkChan <- gameDate.ScheduledGames[0].Link
	}
}

func today() string {
	time := time.Now()
	year, month, day := time.Date()
	return strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
}
