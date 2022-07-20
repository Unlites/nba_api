package stat

import (
	"github.com/Unlites/nba_api/internal/models"
)

type Repository interface {
	GetById(id int64) (*models.Stat, error)
	GetAllByGameId(gameId int64) ([]*models.Stat, error)
	GetAvgByPlayerId(id int64) (*models.AvgByPlayerIdStat, error)
	Create(stat *models.Stat) error
	Update(stat *models.Stat) error
	Delete(id int64) error
}
