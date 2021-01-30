package models

type Team struct {
	Id           int
	Name         string
	EmployerName string
	Player       []Player
	TotalSteps   int
}

type Player struct {
	Id           int
	Name         string
	Steps        int
	EmployerName string
}

type Employer struct {
	Id    int
	Name  string
	Teams []Team
}
