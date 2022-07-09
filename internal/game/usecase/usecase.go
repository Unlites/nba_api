package usecase

import (
	"github.com/Unlites/nba_api/internal/game"
	"github.com/Unlites/nba_api/internal/models"
)

type gameUC struct {
	gameRepo game.Repository
}

func NewGameUseCase(gameRepo game.Repository) game.UseCase {
	return &gameUC{gameRepo: gameRepo}
}

func (uc gameUC) GetById(id int64) (*models.Game, error) {
	return uc.gameRepo.GetById(id)
}

func (uc gameUC) Create(game *models.Game) error {
	return uc.gameRepo.Create(game)
}
