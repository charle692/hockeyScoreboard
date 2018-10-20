package models

// LiveData - Contains live game related data
type LiveData struct {
	Linescore struct {
		CurrentPeriod              int    `json:"currentPeriod"`
		CurrentPeriodOrdinal       string `json:"currentPeriodOrdinal"`
		CurrentPeriodTimeRemaining string `json:"currentPeriodTimeRemaining"`
		Teams                      struct {
			Home struct {
				Team struct {
					ID           int    `json:"id"`
					Name         string `json:"name"`
					Link         string `json:"link"`
					Abbreviation string `json:"abbreviation"`
					TriCode      string `json:"triCode"`
				} `json:"team"`
				Goals        int  `json:"goals"`
				ShotsOnGoal  int  `json:"shotsOnGoal"`
				GoaliePulled bool `json:"goaliePulled"`
				NumSkaters   int  `json:"numSkaters"`
				PowerPlay    bool `json:"powerPlay"`
			} `json:"home"`
			Away struct {
				Team struct {
					ID           int    `json:"id"`
					Name         string `json:"name"`
					Link         string `json:"link"`
					Abbreviation string `json:"abbreviation"`
					TriCode      string `json:"triCode"`
				} `json:"team"`
				Goals        int  `json:"goals"`
				ShotsOnGoal  int  `json:"shotsOnGoal"`
				GoaliePulled bool `json:"goaliePulled"`
				NumSkaters   int  `json:"numSkaters"`
				PowerPlay    bool `json:"powerPlay"`
			} `json:"away"`
		} `json:"teams"`
		PowerPlayStrength string `json:"powerPlayStrength"`
		HasShootout       bool   `json:"hasShootout"`
		IntermissionInfo  struct {
			IntermissionTimeRemaining int  `json:"intermissionTimeRemaining"`
			IntermissionTimeElapsed   int  `json:"intermissionTimeElapsed"`
			InIntermission            bool `json:"inIntermission"`
		} `json:"intermissionInfo"`
		PowerPlayInfo struct {
			SituationTimeRemaining int  `json:"situationTimeRemaining"`
			SituationTimeElapsed   int  `json:"situationTimeElapsed"`
			InSituation            bool `json:"inSituation"`
		} `json:"powerPlayInfo"`
	} `json:"linescore"`
}
