package usecase

import (
	"errors"
	"testing"

	"github.com/Unlites/nba_api/internal/models"
	mock_player "github.com/Unlites/nba_api/internal/player/mocks"
	mock_team "github.com/Unlites/nba_api/internal/team/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name          string
	teamId        int64
	teamInput     *models.Team
	mockBehavior  func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team)
	expected      *models.Team
	isErrorExists bool
}

func TestUseCase_GetById(t *testing.T) {
	testTable := []testCase{
		{
			name:   "Success",
			teamId: 1,
			mockBehavior: func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team) {
				teamRepo.EXPECT().GetById(teamId).Return(&models.Team{
					Id:      1,
					Name:    "Team Name",
					Players: nil,
				}, nil)
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
			expected: &models.Team{
				Id:   1,
				Name: "Team Name",
				Players: []*models.Player{
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
			},
			isErrorExists: false,
		},
		{
			name:   "Error while getting team",
			teamId: 1,
			mockBehavior: func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team) {
				teamRepo.EXPECT().GetById(teamId).Return(nil, errors.New("some error"))
			},
			expected:      nil,
			isErrorExists: true,
		},
		{
			name:   "Error while getting players",
			teamId: 1,
			mockBehavior: func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team) {
				teamRepo.EXPECT().GetById(teamId).Return(&models.Team{
					Id:      1,
					Name:    "Team Name",
					Players: nil,
				}, nil)
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

			teamRepo := mock_team.NewMockRepository(ctrl)
			playerRepo := mock_player.NewMockRepository(ctrl)
			test.mockBehavior(teamRepo, playerRepo, test.teamId, test.teamInput)

			uc := NewTeamUseCase(teamRepo, playerRepo)
			team, err := uc.GetById(test.teamId)
			if test.isErrorExists {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, team)
			}
		})
	}
}

func TestUseCase_Create(t *testing.T) {
	testTable := []testCase{
		{
			name: "Success",
			teamInput: &models.Team{
				Id:      1,
				Name:    "Team Name",
				Players: nil,
			},
			mockBehavior: func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team) {
				teamRepo.EXPECT().Create(teamInput).Return(nil)
			},
			isErrorExists: false,
		},
		{
			name: "Error",
			teamInput: &models.Team{
				Id:      1,
				Name:    "Team Name",
				Players: nil,
			},
			mockBehavior: func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team) {
				teamRepo.EXPECT().Create(teamInput).Return(errors.New("some error"))
			},
			isErrorExists: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			teamRepo := mock_team.NewMockRepository(ctrl)
			playerRepo := mock_player.NewMockRepository(ctrl)
			test.mockBehavior(teamRepo, playerRepo, test.teamId, test.teamInput)

			uc := NewTeamUseCase(teamRepo, playerRepo)
			err := uc.Create(test.teamInput)
			if test.isErrorExists {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUseCase_Update(t *testing.T) {
	testTable := []testCase{
		{
			name: "Success",
			teamInput: &models.Team{
				Id:      1,
				Name:    "Team Name",
				Players: nil,
			},
			mockBehavior: func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team) {
				teamRepo.EXPECT().GetById(teamInput.Id).Return(teamInput, nil)
				teamRepo.EXPECT().Update(teamInput).Return(nil)
			},
			isErrorExists: false,
		},
		{
			name: "Error while getting team",
			teamInput: &models.Team{
				Id:      1,
				Name:    "Team Name",
				Players: nil,
			},
			mockBehavior: func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team) {
				teamRepo.EXPECT().GetById(teamInput.Id).Return(nil, errors.New("some error"))
			},
			isErrorExists: true,
		},
		{
			name: "Error while updating",
			teamInput: &models.Team{
				Id:      1,
				Name:    "Team Name",
				Players: nil,
			},
			mockBehavior: func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team) {
				teamRepo.EXPECT().GetById(teamInput.Id).Return(teamInput, nil)
				teamRepo.EXPECT().Update(teamInput).Return(errors.New("some error"))
			},
			isErrorExists: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			teamRepo := mock_team.NewMockRepository(ctrl)
			playerRepo := mock_player.NewMockRepository(ctrl)
			test.mockBehavior(teamRepo, playerRepo, test.teamId, test.teamInput)

			uc := NewTeamUseCase(teamRepo, playerRepo)
			err := uc.Update(test.teamInput)
			if test.isErrorExists {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUseCase_Delete(t *testing.T) {
	testTable := []testCase{
		{
			name:   "Success",
			teamId: 1,
			mockBehavior: func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team) {
				teamRepo.EXPECT().GetById(teamId).Return(&models.Team{
					Id:      1,
					Name:    "Team Name",
					Players: nil,
				}, nil)
				teamRepo.EXPECT().Delete(teamId).Return(nil)
			},
			isErrorExists: false,
		},
		{
			name:   "Error while getting team",
			teamId: 1,
			mockBehavior: func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team) {
				teamRepo.EXPECT().GetById(teamId).Return(nil, errors.New("some error"))
			},
			isErrorExists: true,
		},
		{
			name:   "Error while deleting",
			teamId: 1,
			mockBehavior: func(teamRepo *mock_team.MockRepository, playerRepo *mock_player.MockRepository, teamId int64, teamInput *models.Team) {
				teamRepo.EXPECT().GetById(teamId).Return(&models.Team{
					Id:      1,
					Name:    "Team Name",
					Players: nil,
				}, nil)
				teamRepo.EXPECT().Delete(teamId).Return(errors.New("some error"))
			},
			isErrorExists: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			teamRepo := mock_team.NewMockRepository(ctrl)
			playerRepo := mock_player.NewMockRepository(ctrl)
			test.mockBehavior(teamRepo, playerRepo, test.teamId, test.teamInput)

			uc := NewTeamUseCase(teamRepo, playerRepo)
			err := uc.Delete(test.teamId)
			if test.isErrorExists {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
