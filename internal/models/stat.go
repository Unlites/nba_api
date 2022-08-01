package models

type Stat struct {
	Id       int64  `json:"id" db:"id"`
	GameId   int64  `json:"game_id" binding:"required,gt=0" db:"game_id"`
	PlayerId int64  `json:"player_id" binding:"required,gt=0" db:"player_id"`
	Points   string `json:"points" binding:"required,number,gte=0" db:"points"`
	Rebounds string `json:"rebounds" binding:"required,number,gte=0" db:"rebounds"`
	Assists  string `json:"assists" binding:"required,number,gte=0" db:"assists"`
}

type AvgByPlayerIdStat struct {
	AvgPoints   float32 `json:"avg_points" db:"avg_points"`
	AvgRebounds float32 `json:"avg_rebounds" db:"avg_rebounds"`
	AvgAssists  float32 `json:"avg_assists" db:"avg_assists"`
}
