package models

type Game struct {
	Id            int64   `json:"id" db:"id"`
	HomeTeamId    int64   `json:"home_team_id" binding:"required" db:"home_team_id"`
	VisitorTeamId int64   `json:"visitor_team_id" binding:"required" db:"visitor_team_id"`
	Score         string  `json:"score" binding:"required" db:"score"`
	WonTeamId     int64   `json:"won_team_id" binding:"required" db:"won_team"`
	Stats         []*Stat `json:"stats" binding:"required"`
}
