package models

type Player struct {
	Id     int64  `json:"id" db:"id"`
	Name   string `json:"name" binding:"required" db:"name"`
	TeamId int64  `json:"team_id" binding:"required" db:"team_id"`
}
