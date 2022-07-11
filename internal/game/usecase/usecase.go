package usecase

import (
	"github.com/Unlites/nba_api/internal/game"
	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/stat"
)

type gameUC struct {
	gameRepo game.Repository
	statRepo stat.Repository
}

func NewGameUseCase(gameRepo game.Repository) game.UseCase {
	return &gameUC{gameRepo: gameRepo}
}

func (uc gameUC) GetById(id int64) (*models.Game, error) {

	game, err := uc.gameRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	stats, err := uc.statRepo.GetAllByGameId(game.Id)
	if err != nil {
		return nil, err
	}

	game.Stats = stats

	return game, nil
}

func (uc gameUC) Create(game *models.Game) error {
	gameId, err := uc.gameRepo.Create(game)
	if err != nil {
		return err
	}

	for _, stat := range game.Stats {
		stat.GameId = gameId
		err = uc.statRepo.Create(stat)
		if err != nil {
			return err
		}
	}

	return err
}
