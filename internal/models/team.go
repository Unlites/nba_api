package models

type Team struct {
	Id   int64  `db:"id"`
	Name string `json:"name" binding:"required" db:"name"`
}
