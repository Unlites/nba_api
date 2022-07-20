package usecase

import (
	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/player"
	"github.com/Unlites/nba_api/internal/team"
)

type teamUC struct {
	teamRepo   team.Repository
	playerRepo player.Repository
}

func NewTeamUseCase(teamRepo team.Repository, playerRepo player.Repository) team.UseCase {
	return &teamUC{teamRepo: teamRepo, playerRepo: playerRepo}
}

func (uc teamUC) GetById(id int64) (*models.Team, error) {
	team, err := uc.teamRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	players, err := uc.playerRepo.GetAllByTeamId(team.Id)
	if err != nil {
		return nil, err
	}

	team.Players = players

	return team, nil
}

func (uc teamUC) Create(team *models.Team) error {
	return uc.teamRepo.Create(team)
}

func (uc teamUC) Update(team *models.Team) error {
	_, err := uc.teamRepo.GetById(team.Id)
	if err != nil {
		return err
	}

	return uc.teamRepo.Update(team)
}

func (uc teamUC) Delete(id int64) error {
	_, err := uc.teamRepo.GetById(id)
	if err != nil {
		return err
	}
	return uc.teamRepo.Delete(id)
}
