package stat

import (
	"github.com/Unlites/nba_api/internal/models"
)

type UseCase interface {
	GetById(id int64) (*models.Stat, error)
	Create(stat *models.Stat) error
	Update(stat *models.Stat) error
	Delete(id int64) error
}
