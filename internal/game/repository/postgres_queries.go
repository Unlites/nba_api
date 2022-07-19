package repository

const (
	gamesTable = "games"

	insertGameQuery     = "INSERT INTO %s (home_team_id, visitor_team_id, score, won_team_id) values ($1, $2, $3, $4) RETURNING id"
	selectGameByIdQuery = "SELECT * FROM %s WHERE id = $1"
	updateGameQuery     = "UPDATE %s SET home_team_id = $2, visitor_team_id = $3, score = $4, won_team_id = $5 WHERE id = $1"
	deleteGameQuery     = "DELETE FROM %s WHERE id = $1"
)
