package footballdata

type parsedTeam struct {
	TLA string `json:"tla"`
}

type parsedMatches struct {
	Matches []struct {
		ID       int        `json:"id"`
		UTCDate  string     `json:"utcDate"`
		Status   string     `json:"status"`
		Stage    string     `json:"stage"`
		HomeTeam parsedTeam `json:"homeTeam"`
		AwayTeam parsedTeam `json:"awayTeam"`
		Score    struct {
			Winner   string `json:"winner"`
			Duration string `json:"duration"`
			FullTime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"fullTime"`
		} `json:"score"`
	}
}

type parsedStandings struct {
	Groups []struct {
		Group string `json:"group"`
		Table []struct {
			Position       int `json:"position"`
			Team           parsedTeam
			PlayedGames    int `json:"playedGames"`
			Won            int `json:"won"`
			Draw           int `json:"draw"`
			Lost           int `json:"lost"`
			GoalsFor       int `json:"goalsFor"`
			GoalsAgainst   int `json:"goalsAgainst"`
			GoalDifference int `json:"goalDifference"`
			Points         int `json:"points"`
		}
	} `json:"standings"`
}
