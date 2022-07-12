package repository

const (
	statsTable = "stats"

	insertStatQuery         = "INSERT INTO %s (game_id, player_id, points, rebounds, assists) values ($1, $2, $3, $4, $5)"
	selectStatByIdQuery     = "SELECT * FROM %s WHERE id = $1"
	selectStatByGameIdQuery = "SELECT * FROM %s WHERE game_id = $1"
)
