package repository

import (
	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/stat"
	"github.com/jmoiron/sqlx"
)

type statRepo struct {
	db *sqlx.DB
}

func NewStatRepo(db *sqlx.DB) stat.Repository {
	return &statRepo{db: db}
}

func (r *statRepo) GetById(id int64) (*models.Stat, error) {
	return nil, nil
}

func (r *statRepo) Create(stat *models.Stat) error {
	return nil
}
