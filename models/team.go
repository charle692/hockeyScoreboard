package models

import (
	"github.com/charle692/hockeyScoreBoard/database"
	"github.com/jinzhu/gorm"
)

// Team - Contains team data
type Team struct {
	gorm.Model
	Selected bool
	Name     string `json:"name"`
	NHLId    int    `json:"id"`
}

func getSelectedTeamName() string {
	return getSelectedTeam().Name
}

func getSelectedTeam() *Team {
	selectedTeam := &Team{}
	database.Conn.Where(&Team{Selected: true}).First(selectedTeam)
	return selectedTeam
}

func getTeams() *[]Team {
	teams := &[]Team{}
	database.Conn.Order("name asc").Find(teams)
	return teams
}

func getTeamByName(teamName string) *Team {
	team := &Team{}
	database.Conn.Where(&Team{Name: teamName}).First(team)
	return team
}
