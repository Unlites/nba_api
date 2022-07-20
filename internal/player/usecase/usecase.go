package usecase

import (
	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/player"
)

type playerUC struct {
	playerRepo player.Repository
}

func NewPlayerUseCase(playerRepo player.Repository) player.UseCase {
	return &playerUC{playerRepo: playerRepo}
}

func (uc playerUC) GetById(id int64) (*models.Player, error) {
	return uc.playerRepo.GetById(id)
}

func (uc playerUC) GetAllByTeamId(teamId int64) ([]*models.Player, error) {
	return uc.playerRepo.GetAllByTeamId(teamId)
}

func (uc playerUC) Create(player *models.Player) error {
	return uc.playerRepo.Create(player)
}

func (uc playerUC) Update(player *models.Player) error {
	_, err := uc.playerRepo.GetById(player.Id)
	if err != nil {
		return err
	}

	return uc.playerRepo.Update(player)
}

func (uc playerUC) Delete(id int64) error {
	_, err := uc.playerRepo.GetById(id)
	if err != nil {
		return err
	}

	return uc.playerRepo.Delete(id)
}
