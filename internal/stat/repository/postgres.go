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
	stat := &models.Stat{}
	query := fmt.Sprintf(selectStatByIdQuery, statsTable)
	err := r.db.Get(stat, query, id)

	return stat, err
}

func (r *statRepo) GetAllByGameId(gameId int64) ([]*models.Stat, error) {
	stats := make([]*models.Stat, 0)
	query := fmt.Sprintf(selectStatsByGameIdQuery, statsTable)
	err := r.db.Select(&stats, query, gameId)

	return stats, err
}

func (r *statRepo) GetAvgByPlayerId(id int64) (*models.AvgByPlayerIdStat, error) {
	avgStat := &models.AvgByPlayerIdStat{}
	query := fmt.Sprintf(selectAvgStatByPlayerIdQuery, statsTable)
	err := r.db.Get(avgStat, query, id)

	return avgStat, err
}

func (r *statRepo) Create(stat *models.Stat) error {
	query := fmt.Sprintf(insertStatQuery, statsTable)
	_, err := r.db.Exec(query, stat.GameId, stat.PlayerId, stat.Points, stat.Rebounds, stat.Assists)

	return err
}

func (r *statRepo) Update(stat *models.Stat) error {
	query := fmt.Sprintf(updateStatQuery, statsTable)
	_, err := r.db.Exec(query, stat.Id, stat.GameId, stat.PlayerId, stat.Points, stat.Rebounds, stat.Assists)

	return err
}

func (r *statRepo) Delete(id int64) error {
	query := fmt.Sprintf(deleteStatQuery, statsTable)
	_, err := r.db.Exec(query, id)

	return err
}
