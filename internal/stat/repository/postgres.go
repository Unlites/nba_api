package repository

import (
	"fmt"

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
	var stat *models.Stat
	query := fmt.Sprintf(selectStatByIdQuery, statsTable)
	err := r.db.Get(&stat, query, id)

	return stat, err
}

func (r *statRepo) GetAllByGameId(gameId int64) ([]*models.Stat, error) {
	var stats []*models.Stat
	query := fmt.Sprintf(selectStatByGameIdQuery, statsTable)
	err := r.db.Select(&stats, query, gameId)

	return stats, err
}

func (r *statRepo) Create(stat *models.Stat) error {
	query := fmt.Sprintf(insertStatQuery, statsTable)
	_, err := r.db.Exec(query, stat.GameId, stat.PlayerId, stat.Points, stat.Rebounds, stat.Assists)

	return err
}
