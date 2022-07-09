package player

import (
	"github.com/Unlites/nba_api/internal/models"
)

type UseCase interface {
	GetById(id int64) (*models.Player, error)
	Create(stat *models.Player) error
}
