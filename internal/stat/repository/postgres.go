package repository

import (
	"fmt"

	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/stat"
	"github.com/jmoiron/sqlx"
)

const statsTable = "stats"

type statRepo struct {
	db *sqlx.DB
}

func NewStatRepo(db *sqlx.DB) stat.Repository {
	return &statRepo{db: db}
}

func (r *statRepo) GetById(id int64) (*models.Stat, error) {
	return nil, nil
}

func (r *statRepo) GetByGameId(gameId int64) ([]*models.Stat, error) {
	return nil, nil
}

func (r *statRepo) Create(stat *models.Stat) error {
	query := fmt.Sprintf(insertStatQuery, statsTable)
	_, err := r.db.Exec(query, stat.GameId, stat.PlayerId, stat.Points, stat.Rebounds, stat.Assists)

	return err
}
