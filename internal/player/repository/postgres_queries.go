package repository

const (
	playersTable = "players"

	insertPlayerQuery       = "INSERT INTO %s (name, team_id) VALUES ($1, $2)"
	selectPlayerByIdQuery   = "SELECT * FROM %s WHERE id = $1"
	selectStatByTeamIdQuery = "SELECT * FROM %s WHERE team_id = $1"
	updatePlayerQuery       = "UPDATE %s SET name = $2, team_id = $3 WHERE id = $1"
	deletePlayerQuery       = "DELETE FROM %s WHERE id = $1"
)
