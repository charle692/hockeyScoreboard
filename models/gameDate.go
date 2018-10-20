package models

// GameDate - Contains a team's games for the day
type GameDate struct {
	Date           string          `json:"date"`
	TotalItems     int             `json:"totalItems"`
	TotalEvents    int             `json:"totalEvents"`
	TotalGames     int             `json:"totalGames"`
	TotalMatches   int             `json:"totalMatches"`
	ScheduledGames []ScheduledGame `json:"games"`
	Events         []interface{}   `json:"events"`
	Matches        []interface{}   `json:"matches"`
}
