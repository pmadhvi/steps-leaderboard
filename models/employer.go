package model

//Ex:
//Employer = Employer{ID: 1, Name: TCS, Teams: [{ID: 1, Name: ABC, Players: [{ID: 1, Name: ABC, Steps: 200}, {ID: 2, Name: XYZ, Steps: 300}], TotalSteps: 2000}, ]}
//Employer = EMployer{ID: 2, Name: Mahindra}

type Employer struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Teams []Team `json:"team"`
}

type Employers []Employer
