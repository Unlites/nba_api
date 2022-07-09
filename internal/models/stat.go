package models

type Stat struct {
	Id       int64 `json:"id" db:"id"`
	GameId   int64 `json:"game_id" binding:"required" db:"game_id"`
	PlayerId int64 `json:"player_id" binding:"required" db:"player_id"`
	Points   int64 `json:"points" binding:"required" db:"points"`
	Rebounds int64 `json:"rebounds" binding:"required" db:"rebounds"`
	Assists  int64 `json:"assists" binding:"required" db:"assists"`
}
