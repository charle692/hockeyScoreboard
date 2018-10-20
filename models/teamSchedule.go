package models

// TeamSchedule - Contains a team's schedule for the day
type TeamSchedule struct {
	TotalEvents  int        `json:"totalEvents"`
	TotalGames   int        `json:"totalGames"`
	TotalMatches int        `json:"totalMatches"`
	Wait         int        `json:"wait"`
	GameDates    []GameDate `json:"dates"`
}
