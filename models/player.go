package model

//Ex:
//Player = Player{ID: 1, Name: ABC, Steps: 200, EmpName: TCS}
//Player = Player{ID: 2, Name: DEF, Steps: 500, EmpName: TCS}
//Player = Player{ID: 3, Name: XYZ, Steps: 300, EmpName: Mahindra}
//Player = Player{ID: 4, Name: EFG, Steps: 500, EmpName: Mahindra}

type Player struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Steps    int64  `json:"steps"`
	EmpName  string `json:"emp_name"`
	TeamName string `json:"team_name"`
}

type Players []Player
