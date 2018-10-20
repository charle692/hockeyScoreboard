package models

type GameData struct {
	Status struct {
		AbstractGameState string `json:"abstractGameState"`
		CodedGameState    string `json:"codedGameState"`
		DetailedState     string `json:"detailedState"`
		StatusCode        string `json:"statusCode"`
		StartTimeTBD      bool   `json:"startTimeTBD"`
	} `json:"status"`
	Teams struct {
		Away struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			Link            string `json:"link"`
			Abbreviation    string `json:"abbreviation"`
			TriCode         string `json:"triCode"`
			TeamName        string `json:"teamName"`
			LocationName    string `json:"locationName"`
			FirstYearOfPlay string `json:"firstYearOfPlay"`
			ShortName       string `json:"shortName"`
			OfficialSiteURL string `json:"officialSiteUrl"`
			FranchiseID     int    `json:"franchiseId"`
			Active          bool   `json:"active"`
		} `json:"away"`
		Home struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			Link            string `json:"link"`
			Abbreviation    string `json:"abbreviation"`
			TriCode         string `json:"triCode"`
			TeamName        string `json:"teamName"`
			LocationName    string `json:"locationName"`
			FirstYearOfPlay string `json:"firstYearOfPlay"`
			ShortName       string `json:"shortName"`
			OfficialSiteURL string `json:"officialSiteUrl"`
			FranchiseID     int    `json:"franchiseId"`
			Active          bool   `json:"active"`
		} `json:"home"`
	} `json:"teams"`
}
