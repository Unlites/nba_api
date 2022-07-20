package usecase

import (
	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/stat"
)

type statUC struct {
	statRepo stat.Repository
}

func NewStatUseCase(statRepo stat.Repository) stat.UseCase {
	return &statUC{statRepo: statRepo}
}

func (uc statUC) GetById(id int64) (*models.Stat, error) {
	return uc.statRepo.GetById(id)
}

func (uc statUC) GetAvgByPlayerId(id int64) (*models.AvgByPlayerIdStat, error) {
	return uc.statRepo.GetAvgByPlayerId(id)
}

func (uc statUC) Create(stat *models.Stat) error {
	return uc.statRepo.Create(stat)
}

func (uc statUC) Update(stat *models.Stat) error {
	_, err := uc.statRepo.GetById(stat.Id)
	if err != nil {
		return err
	}

	return uc.statRepo.Update(stat)
}

func (uc statUC) Delete(id int64) error {
	_, err := uc.statRepo.GetById(id)
	if err != nil {
		return err
	}

	return uc.statRepo.Delete(id)
}
