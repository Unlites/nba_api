package repository

import (
	"fmt"

	"github.com/Unlites/nba_api/internal/game"
	"github.com/Unlites/nba_api/internal/models"
	"github.com/jmoiron/sqlx"
)

type gameRepo struct {
	db *sqlx.DB
}

func NewGameRepo(db *sqlx.DB) game.Repository {
	return &gameRepo{db: db}
}

func (r *gameRepo) GetById(id int64) (*models.Game, error) {
	game := &models.Game{}
	query := fmt.Sprintf(selectGameByIdQuery, gamesTable)
	err := r.db.Get(game, query, id)

	return game, err
}

func (r *gameRepo) Create(game *models.Game) (int64, error) {
	query := fmt.Sprintf(insertGameQuery, gamesTable)
	row := r.db.QueryRow(query, game.HomeTeamId, game.VisitorTeamId, game.Score, game.WonTeamId)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
