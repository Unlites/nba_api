package repository

import (
	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/player"
	"github.com/jmoiron/sqlx"
)

type playerRepo struct {
	db *sqlx.DB
}

func NewPlayerRepo(db *sqlx.DB) player.Repository {
	return &playerRepo{db: db}
}

func (r *playerRepo) GetById(id int64) (*models.Player, error) {
	return nil, nil
}

func (r *playerRepo) Create(stat *models.Player) error {
	return nil
}
