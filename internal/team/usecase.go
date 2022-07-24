package team

import (
	"github.com/Unlites/nba_api/internal/models"
)

type UseCase interface {
	GetById(id int64) (*models.Team, error)
	Create(team *models.Team) error
	Update(team *models.Team) error
	Delete(id int64) error
}
