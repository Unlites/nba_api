package player

import (
	"github.com/Unlites/nba_api/internal/models"
)

type UseCase interface {
	GetById(id int64) (*models.Player, error)
	GetAllByTeamId(teamId int64) ([]*models.Player, error)
	Create(game *models.Player) error
	Update(game *models.Player) error
	Delete(id int64) error
}
