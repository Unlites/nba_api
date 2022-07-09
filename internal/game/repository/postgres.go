package repository

import (
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
	return nil, nil
}

func (r *gameRepo) Create(game *models.Game) error {
	return nil
}
