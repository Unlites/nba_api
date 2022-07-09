package team

import (
	"github.com/Unlites/nba_api/internal/models"
)

type Repository interface {
	GetById(id int64) (*models.Team, error)
	Create(stat *models.Team) error
}
