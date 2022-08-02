package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_game "github.com/Unlites/nba_api/internal/game/mocks"
	"github.com/Unlites/nba_api/internal/models"
	httpErr "github.com/Unlites/nba_api/pkg/http_errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name                 string
	queryParamId         int64
	queryBody            *models.Game
	mockBehavior         func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game)
	expectedStatusCode   int
	expectedResponseBody string
}

type gameHandlerTestParams struct {
	t            *testing.T
	testCase     testCase
	queryMethod  string
	queryURL     string
	isBodyExists bool
}

func gameHandlerTestRun(testParams *gameHandlerTestParams) {
	ctrl := gomock.NewController(testParams.t)
	defer ctrl.Finish()

	uc := mock_game.NewMockUseCase(ctrl)
	testParams.testCase.mockBehavior(uc, testParams.testCase.queryParamId, testParams.testCase.queryBody)

	handler := NewGameHandler(uc)

	router := gin.Default()
	group := router.Group("/api/v1/game")

	RegisterRoutes(group, handler)

	var req *http.Request

	if testParams.isBodyExists {
		body, err := json.Marshal(testParams.testCase.queryBody)
		assert.NoError(testParams.t, err)
		req = httptest.NewRequest(testParams.queryMethod, testParams.queryURL, bytes.NewBuffer(body))
	} else {
		req = httptest.NewRequest(testParams.queryMethod, testParams.queryURL, nil)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(testParams.t, testParams.testCase.expectedStatusCode, w.Code)
	assert.Contains(testParams.t, w.Body.String(), testParams.testCase.expectedResponseBody)
}

func TestHandler_GetById(t *testing.T) {
	testTable := []testCase{
		{
			name:         "Success",
			queryParamId: 1,
			mockBehavior: func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {
				s.EXPECT().GetById(queryParamId).Return(&models.Game{
					Id:            1,
					HomeTeamId:    1,
					VisitorTeamId: 2,
					Score:         "120:100",
					WonTeamId:     1,
					Stats: []*models.Stat{
						{
							Id:       1,
							GameId:   1,
							PlayerId: 1,
							Points:   "30",
							Rebounds: "8",
							Assists:  "6",
						},
						{
							Id:       2,
							GameId:   1,
							PlayerId: 2,
							Points:   "25",
							Rebounds: "5",
							Assists:  "3",
						},
					},
				}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{` +
				`"id":1,` +
				`"home_team_id":1,` +
				`"visitor_team_id":2,` +
				`"score":"120:100",` +
				`"won_team_id":1,` +
				`"stats":[` +
				`{` +
				`"id":1,` +
				`"game_id":1,` +
				`"player_id":1,` +
				`"points":"30",` +
				`"rebounds":"8",` +
				`"assists":"6"` +
				`},` +
				`{` +
				`"id":2,` +
				`"game_id":1,` +
				`"player_id":2,` +
				`"points":"25",` +
				`"rebounds":"5",` +
				`"assists":"3"` +
				`}` +
				`]` +
				`}`,
		},
		{
			name:                 "Not positive id",
			queryParamId:         0,
			mockBehavior:         func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrNotPositiveId,
		},
		{
			name:         "Not found",
			queryParamId: 2,
			mockBehavior: func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {
				s.EXPECT().GetById(queryParamId).Return(nil, errors.New("sql: no rows in result set"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: httpErr.ErrNotFound,
		},
		{
			name:         "Internal error",
			queryParamId: 1,
			mockBehavior: func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {
				s.EXPECT().GetById(queryParamId).Return(nil, errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			gameHandlerTestRun(&gameHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "GET",
				queryURL:     fmt.Sprintf("/api/v1/game/%v", testCase.queryParamId),
				isBodyExists: false,
			})
		})
	}
}

func TestHandler_Create(t *testing.T) {
	testTable := []testCase{
		{
			name: "Success",
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				Score:         "120:100",
				WonTeamId:     1,
				Stats: []*models.Stat{
					{
						GameId:   1,
						PlayerId: 1,
						Points:   "30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						GameId:   1,
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior: func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {
				s.EXPECT().Create(queryBody).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name: "Missed param",
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				WonTeamId:     1,
				Stats: []*models.Stat{
					{
						PlayerId: 1,
						Points:   "30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior:         func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name: "Empty param",
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				Score:         "",
				WonTeamId:     1,
				Stats: []*models.Stat{
					{
						GameId:   1,
						PlayerId: 1,
						Points:   "30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						GameId:   1,
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior:         func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name: "Internal error",
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				Score:         "120:100",
				WonTeamId:     1,
				Stats: []*models.Stat{
					{
						GameId:   1,
						PlayerId: 1,
						Points:   "30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						GameId:   1,
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior: func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {
				s.EXPECT().Create(queryBody).Return(errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
		{
			name: "Invalid won team id",
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				Score:         "120:100",
				WonTeamId:     3,
				Stats: []*models.Stat{
					{
						PlayerId: 1,
						Points:   "30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior:         func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name: "Not positive param",
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				Score:         "120:100",
				WonTeamId:     3,
				Stats: []*models.Stat{
					{
						PlayerId: 1,
						Points:   "-30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior:         func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			gameHandlerTestRun(&gameHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "POST",
				queryURL:     "/api/v1/game/",
				isBodyExists: true,
			})
		})
	}
}

func TestHandler_Update(t *testing.T) {
	testTable := []testCase{
		{
			name:         "Success",
			queryParamId: 1,
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				Score:         "120:100",
				WonTeamId:     1,
				Stats: []*models.Stat{
					{
						GameId:   1,
						PlayerId: 1,
						Points:   "30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						GameId:   1,
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior: func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {
				s.EXPECT().Update(queryBody).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:         "Not positive id",
			queryParamId: 0,
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				Score:         "120:100",
				WonTeamId:     1,
				Stats: []*models.Stat{
					{
						GameId:   1,
						PlayerId: 1,
						Points:   "30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						GameId:   1,
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior:         func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrNotPositiveId,
		},
		{
			name:         "Not found",
			queryParamId: 1,
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				Score:         "120:100",
				WonTeamId:     1,
				Stats: []*models.Stat{
					{
						GameId:   1,
						PlayerId: 1,
						Points:   "30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						GameId:   1,
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior: func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {
				s.EXPECT().Update(queryBody).Return(errors.New("sql: no rows in result set"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: httpErr.ErrNotFound,
		},
		{
			name:         "Missed param",
			queryParamId: 1,
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				WonTeamId:     1,
				Stats: []*models.Stat{
					{
						GameId:   1,
						PlayerId: 1,
						Points:   "30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						GameId:   1,
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior:         func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name:         "Empty param",
			queryParamId: 1,
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				Score:         "",
				WonTeamId:     1,
				Stats: []*models.Stat{
					{
						GameId:   1,
						PlayerId: 1,
						Points:   "30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						GameId:   1,
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior:         func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name:         "Internal error",
			queryParamId: 1,
			queryBody: &models.Game{
				HomeTeamId:    1,
				VisitorTeamId: 2,
				Score:         "120:100",
				WonTeamId:     1,
				Stats: []*models.Stat{
					{
						GameId:   1,
						PlayerId: 1,
						Points:   "30",
						Rebounds: "8",
						Assists:  "6",
					},
					{
						GameId:   1,
						PlayerId: 2,
						Points:   "25",
						Rebounds: "5",
						Assists:  "3",
					},
				},
			},
			mockBehavior: func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {
				s.EXPECT().Update(queryBody).Return(errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			gameHandlerTestRun(&gameHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "PUT",
				queryURL:     fmt.Sprintf("/api/v1/game/%v", testCase.queryParamId),
				isBodyExists: true,
			})
		})
	}
}

func TestHandler_Delete(t *testing.T) {
	testTable := []testCase{
		{
			name:         "Success",
			queryParamId: 1,
			mockBehavior: func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {
				s.EXPECT().Delete(queryParamId).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:                 "Not positive id",
			queryParamId:         0,
			mockBehavior:         func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrNotPositiveId,
		},
		{
			name:         "Not found",
			queryParamId: 2,
			mockBehavior: func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {
				s.EXPECT().Delete(queryParamId).Return(errors.New("sql: no rows in result set"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: httpErr.ErrNotFound,
		},
		{
			name:         "Internal error",
			queryParamId: 1,
			mockBehavior: func(s *mock_game.MockUseCase, queryParamId int64, queryBody *models.Game) {
				s.EXPECT().Delete(queryParamId).Return(errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			gameHandlerTestRun(&gameHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "DELETE",
				queryURL:     fmt.Sprintf("/api/v1/game/%v", testCase.queryParamId),
				isBodyExists: false,
			})
		})
	}
}
