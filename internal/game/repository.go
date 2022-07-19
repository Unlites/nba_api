package game

import (
	"github.com/Unlites/nba_api/internal/models"
)

type Repository interface {
	GetById(id int64) (*models.Game, error)
	Create(game *models.Game) (int64, error)
	Update(game *models.Game) error
	Delete(id int64) error
}
