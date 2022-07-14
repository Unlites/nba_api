package models

type Team struct {
	Id      int64     `json:"id" db:"id"`
	Name    string    `json:"name" binding:"required" db:"name"`
	Players []*Player `json:"players"`
}
