package repository

const (
	teamsTable = "teams"

	insertTeamQuery     = "INSERT INTO %s (name) VALUES ($1)"
	selectTeamByIdQuery = "SELECT * FROM %s WHERE id = $1"
	updateTeamQuery     = "UPDATE %s SET name = $2 WHERE id = $1"
	deleteTeamQuery     = "DELETE FROM %s WHERE id = $1"
)
