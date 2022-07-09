package usecase

import (
	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/team"
)

type teamUC struct {
	teamRepo team.Repository
}

func NewTeamUseCase(teamRepo team.Repository) team.UseCase {
	return &teamUC{teamRepo: teamRepo}
}

func (uc teamUC) GetById(id int64) (*models.Team, error) {
	return uc.teamRepo.GetById(id)
}

func (uc teamUC) Create(team *models.Team) error {
	return uc.teamRepo.Create(team)
}
