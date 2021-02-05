package models

//Ex:
//Team = Team{ID: 1, Name: ABC, Players: [{ID: 1, Name: ABC, Steps: 200}, {ID: 2, Name: XYZ, Steps: 300}], TotalSteps: 2000}
//Team = Team{ID: 2, Name: ABC, Players: [{ID: 2, Name: XYZ, Steps: 300}], TotalSteps: 5000}

type Team struct {
	Id         int64    `json:"id"`
	Name       string   `json:"name"`
	Players    []Player `json:"players"`
	TotalSteps int64    `json:"total_steps"`
}

type Teams []Team
