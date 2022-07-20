package repository

const (
	statsTable = "stats"

	insertStatQuery              = "INSERT INTO %s (game_id, player_id, points, rebounds, assists) values ($1, $2, $3, $4, $5)"
	selectStatByIdQuery          = "SELECT * FROM %s WHERE id = $1"
	selectStatsByGameIdQuery     = "SELECT * FROM %s WHERE game_id = $1"
	selectAvgStatByPlayerIdQuery = `SELECT ROUND(AVG(points), 2) AS avg_points, 
										   ROUND(AVG(rebounds), 2) AS avg_rebounds, 
										   ROUND(AVG(assists), 2) AS avg_assists 
									FROM %s WHERE player_id = $1`
	updateStatQuery = "UPDATE %s SET game_id = $2, player_id = $3, points = $4, rebounds = $5, assists = $6 WHERE id = $1"
	deleteStatQuery = "DELETE FROM %s WHERE id = $1"
)
