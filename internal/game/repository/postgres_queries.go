package repository

const (
	insertGameQuery     = "INSERT INTO %s (home_team_id, visitor_team_id, score, won_team_id) values ($1, $2, $3, $4) RETURNING id"
	selectGameByIdQuery = "SELECT * FROM %s WHERE id = $1"
)
