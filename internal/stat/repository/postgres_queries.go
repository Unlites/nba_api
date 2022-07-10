package repository

const (
	insertStatQuery = "INSERT INTO %s (game_id, player_id, points, rebounds, assists) values ($1, $2, $3, $4, $5)"
)
