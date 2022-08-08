package usecase

import (
	"errors"
	"testing"

	"github.com/Unlites/nba_api/internal/models"
	mock_player "github.com/Unlites/nba_api/internal/player/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type defaultCRUDtestCase struct {
	name          string
	playerId      int64
	playerInput   *models.Player
	mockBehavior  func(playerRepo *mock_player.MockRepository, playerId int64, playerInput *models.Player)
	expected      *models.Player
	isErrorExists bool
}

func TestUseCase_GetById(t *testing.T) {
	testTable := []defaultCRUDtestCase{
		{
			name:     "Success",
			playerId: 1,
			mockBehavior: func(playerRepo *mock_player.MockRepository, playerId int64, playerInput *models.Player) {
				playerRepo.EXPECT().GetById(playerId).Return(&models.Player{
					Id:     1,
					Name:   "Player Name",
					TeamId: 1,
				}, nil)
			},
			expected: &models.Player{
				Id:     1,
				Name:   "Player Name",
				TeamId: 1,
			},
			isErrorExists: false,
		},
		{
			name:     "Error",
			playerId: 1,
			mockBehavior: func(playerRepo *mock_player.MockRepository, playerId int64, playerInput *models.Player) {
				playerRepo.EXPECT().GetById(playerId).Return(nil, errors.New("some error"))
			},
			expected:      nil,
			isErrorExists: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			playerRepo := mock_player.NewMockRepository(ctrl)
			test.mockBehavior(playerRepo, test.playerId, test.playerInput)

			uc := NewPlayerUseCase(playerRepo)
			player, err := uc.GetById(test.playerId)
			if test.isErrorExists {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, player)
			}
		})
	}
}

func TestUseCase_GetAllByTeamId(t *testing.T) {
	testTable := []struct {
		name          string
		teamId        int64
		mockBehavior  func(playerRepo *mock_player.MockRepository, teamId int64)
		expected      []*models.Player
		isErrorExists bool
	}{
		{
			name:   "Success",
			teamId: 1,
			mockBehavior: func(playerRepo *mock_player.MockRepository, teamId int64) {
				playerRepo.EXPECT().GetAllByTeamId(teamId).Return([]*models.Player{
					{
						Id:     1,
						Name:   "Player Name",
						TeamId: 1,
					},
					{
						Id:     2,
						Name:   "Player Name",
						TeamId: 1,
					},
				}, nil)
			},
			expected: []*models.Player{
				{
					Id:     1,
					Name:   "Player Name",
					TeamId: 1,
				},
				{
					Id:     2,
					Name:   "Player Name",
					TeamId: 1,
				},
			},
			isErrorExists: false,
		},
		{
			name:   "Error",
			teamId: 1,
			mockBehavior: func(playerRepo *mock_player.MockRepository, teamId int64) {
				playerRepo.EXPECT().GetAllByTeamId(teamId).Return(nil, errors.New("some error"))
			},
			expected:      nil,
			isErrorExists: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			playerRepo := mock_player.NewMockRepository(ctrl)
			test.mockBehavior(playerRepo, test.teamId)

			uc := NewPlayerUseCase(playerRepo)
			player, err := uc.GetAllByTeamId(test.teamId)
			if test.isErrorExists {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, player)
			}
		})
	}
}

func TestUseCase_Create(t *testing.T) {
	testTable := []defaultCRUDtestCase{
		{
			name: "Success",
			playerInput: &models.Player{
				Id:     1,
				Name:   "Player Name",
				TeamId: 1,
			},
			mockBehavior: func(playerRepo *mock_player.MockRepository, playerId int64, playerInput *models.Player) {
				playerRepo.EXPECT().Create(playerInput).Return(nil)
			},
			isErrorExists: false,
		},
		{
			name: "Error",
			playerInput: &models.Player{
				Id:     1,
				Name:   "Player Name",
				TeamId: 1,
			},
			mockBehavior: func(playerRepo *mock_player.MockRepository, playerId int64, playerInput *models.Player) {
				playerRepo.EXPECT().Create(playerInput).Return(errors.New("some error"))
			},
			isErrorExists: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			playerRepo := mock_player.NewMockRepository(ctrl)
			test.mockBehavior(playerRepo, test.playerId, test.playerInput)

			uc := NewPlayerUseCase(playerRepo)
			err := uc.Create(test.playerInput)
			if test.isErrorExists {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUseCase_Update(t *testing.T) {
	testTable := []defaultCRUDtestCase{
		{
			name: "Success",
			playerInput: &models.Player{
				Id:     1,
				Name:   "Player Name",
				TeamId: 1,
			},
			mockBehavior: func(playerRepo *mock_player.MockRepository, playerId int64, playerInput *models.Player) {
				playerRepo.EXPECT().GetById(playerInput.Id).Return(&models.Player{
					Id:     1,
					Name:   "Team Name",
					TeamId: 1,
				}, nil)
				playerRepo.EXPECT().Update(playerInput).Return(nil)
			},
			isErrorExists: false,
		},
		{
			name: "Error while getting player",
			playerInput: &models.Player{
				Id:     1,
				Name:   "Player Name",
				TeamId: 1,
			},
			mockBehavior: func(playerRepo *mock_player.MockRepository, playerId int64, playerInput *models.Player) {
				playerRepo.EXPECT().GetById(playerInput.Id).Return(nil, errors.New("some error"))
			},
			isErrorExists: true,
		},
		{
			name: "Error while updating player",
			playerInput: &models.Player{
				Id:     1,
				Name:   "Player Name",
				TeamId: 1,
			},
			mockBehavior: func(playerRepo *mock_player.MockRepository, playerId int64, playerInput *models.Player) {
				playerRepo.EXPECT().GetById(playerInput.Id).Return(&models.Player{
					Id:     1,
					Name:   "Team Name",
					TeamId: 1,
				}, nil)
				playerRepo.EXPECT().Update(playerInput).Return(errors.New("some error"))
			},
			isErrorExists: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			playerRepo := mock_player.NewMockRepository(ctrl)
			test.mockBehavior(playerRepo, test.playerId, test.playerInput)

			uc := NewPlayerUseCase(playerRepo)
			err := uc.Update(test.playerInput)
			if test.isErrorExists {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUseCase_Delete(t *testing.T) {
	testTable := []defaultCRUDtestCase{
		{
			name:     "Success",
			playerId: 1,
			mockBehavior: func(playerRepo *mock_player.MockRepository, playerId int64, playerInput *models.Player) {
				playerRepo.EXPECT().GetById(playerId).Return(&models.Player{
					Id:     1,
					Name:   "Team Name",
					TeamId: 1,
				}, nil)
				playerRepo.EXPECT().Delete(playerId).Return(nil)
			},
			isErrorExists: false,
		},
		{
			name:     "Error while getting player",
			playerId: 1,
			mockBehavior: func(playerRepo *mock_player.MockRepository, playerId int64, playerInput *models.Player) {
				playerRepo.EXPECT().GetById(playerId).Return(nil, errors.New("some error"))
			},
			isErrorExists: true,
		},
		{
			name:     "Error while deleting player",
			playerId: 1,
			mockBehavior: func(playerRepo *mock_player.MockRepository, playerId int64, playerInput *models.Player) {
				playerRepo.EXPECT().GetById(playerId).Return(&models.Player{
					Id:     1,
					Name:   "Team Name",
					TeamId: 1,
				}, nil)
				playerRepo.EXPECT().Delete(playerId).Return(errors.New("some error"))
			},
			isErrorExists: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			playerRepo := mock_player.NewMockRepository(ctrl)
			test.mockBehavior(playerRepo, test.playerId, test.playerInput)

			uc := NewPlayerUseCase(playerRepo)
			err := uc.Delete(test.playerId)
			if test.isErrorExists {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
