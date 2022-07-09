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

func (uc playerUC) Create(player *models.Player) error {
	return uc.playerRepo.Create(player)
}
