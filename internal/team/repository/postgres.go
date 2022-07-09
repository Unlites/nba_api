package repository

import (
	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/team"
	"github.com/jmoiron/sqlx"
)

type teamRepo struct {
	db *sqlx.DB
}

func NewTeamRepo(db *sqlx.DB) team.Repository {
	return &teamRepo{db: db}
}

func (r *teamRepo) GetById(id int64) (*models.Team, error) {
	return nil, nil
}

func (r *teamRepo) Create(stat *models.Team) error {
	return nil
}
