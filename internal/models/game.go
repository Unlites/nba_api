package models

type Game struct {
	Id            int64   `json:"id" db:"id"`
	HomeTeamId    int64   `json:"home_team_id" binding:"required,gt=0" db:"home_team_id"`
	VisitorTeamId int64   `json:"visitor_team_id" binding:"required,gt=0" db:"visitor_team_id"`
	Score         string  `json:"score" binding:"required" db:"score"`
	WonTeamId     int64   `json:"won_team_id" binding:"required,eqfield=HomeTeamId|eqfield=VisitorTeamId" db:"won_team_id"`
	Stats         []*Stat `json:"stats" binding:"dive"`
}
