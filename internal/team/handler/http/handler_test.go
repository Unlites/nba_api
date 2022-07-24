package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Unlites/nba_api/internal/models"
	mock_team "github.com/Unlites/nba_api/internal/team/mocks"
	httpErr "github.com/Unlites/nba_api/pkg/http_errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name                 string
	queryParamId         int64
	queryBody            *models.Team
	mockBehavior         func(s *mock_team.MockUseCase, queryParamId int64, input *models.Team)
	expectedStatusCode   int
	expectedResponseBody string
}

type teamHandlerTestParams struct {
	t            *testing.T
	testCase     testCase
	queryMethod  string
	queryURL     string
	isBodyExists bool
}

func teamHandlerTestRun(testParams *teamHandlerTestParams) {
	ctrl := gomock.NewController(testParams.t)
	defer ctrl.Finish()

	uc := mock_team.NewMockUseCase(ctrl)
	testParams.testCase.mockBehavior(uc, testParams.testCase.queryParamId, testParams.testCase.queryBody)

	handler := NewTeamHandler(uc)

	router := gin.Default()
	group := router.Group("/api/v1/team")

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
			mockBehavior: func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {
				s.EXPECT().GetById(queryParamId).Return(&models.Team{
					Id:      1,
					Name:    "Team Name",
					Players: nil,
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"name":"Team Name","players":null}`,
		},
		{
			name:                 "Not positive id",
			queryParamId:         0,
			mockBehavior:         func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrNotPositiveId,
		},
		{
			name:         "Not found",
			queryParamId: 2,
			mockBehavior: func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {
				s.EXPECT().GetById(queryParamId).Return(nil, errors.New("sql: no rows in result set"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: httpErr.ErrNotFound,
		},
		{
			name:         "Internal error",
			queryParamId: 1,
			mockBehavior: func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {
				s.EXPECT().GetById(queryParamId).Return(nil, errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			teamHandlerTestRun(&teamHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "GET",
				queryURL:     fmt.Sprintf("/api/v1/team/%v", testCase.queryParamId),
				isBodyExists: false,
			})
		})
	}
}

func TestHandler_Create(t *testing.T) {
	testTable := []testCase{
		{
			name: "Success",
			queryBody: &models.Team{
				Name: "Team Name",
			},
			mockBehavior: func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {
				s.EXPECT().Create(queryBody).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:                 "Missed param",
			queryBody:            &models.Team{},
			mockBehavior:         func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name: "Empty param",
			queryBody: &models.Team{
				Name: "",
			},
			mockBehavior:         func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name: "Internal error",
			queryBody: &models.Team{
				Name: "Team Name",
			},
			mockBehavior: func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {
				s.EXPECT().Create(queryBody).Return(errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			teamHandlerTestRun(&teamHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "POST",
				queryURL:     "/api/v1/team/",
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
			queryBody: &models.Team{
				Name: "Team Name",
			},
			mockBehavior: func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {
				s.EXPECT().Update(queryBody).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:         "Not positive id",
			queryParamId: 0,
			queryBody: &models.Team{
				Name: "Team Name",
			},
			mockBehavior:         func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrNotPositiveId,
		},
		{
			name:         "Not found",
			queryParamId: 1,
			queryBody: &models.Team{
				Name: "Team Name",
			},
			mockBehavior: func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {
				s.EXPECT().Update(queryBody).Return(errors.New("sql: no rows in result set"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: httpErr.ErrNotFound,
		},
		{
			name:                 "Missed param",
			queryParamId:         1,
			queryBody:            &models.Team{},
			mockBehavior:         func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name:         "Empty param",
			queryParamId: 1,
			queryBody: &models.Team{
				Name: "",
			},
			mockBehavior:         func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name:         "Internal error",
			queryParamId: 1,
			queryBody: &models.Team{
				Name: "Team Name",
			},
			mockBehavior: func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {
				s.EXPECT().Update(queryBody).Return(errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			teamHandlerTestRun(&teamHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "PUT",
				queryURL:     fmt.Sprintf("/api/v1/team/%v", testCase.queryParamId),
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
			mockBehavior: func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {
				s.EXPECT().Delete(queryParamId).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:                 "Not positive id",
			queryParamId:         0,
			mockBehavior:         func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrNotPositiveId,
		},
		{
			name:         "Not found",
			queryParamId: 2,
			mockBehavior: func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {
				s.EXPECT().Delete(queryParamId).Return(errors.New("sql: no rows in result set"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: httpErr.ErrNotFound,
		},
		{
			name:         "Internal error",
			queryParamId: 1,
			queryBody: &models.Team{
				Name: "Team Name",
			},
			mockBehavior: func(s *mock_team.MockUseCase, queryParamId int64, queryBody *models.Team) {
				s.EXPECT().Delete(queryParamId).Return(errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			teamHandlerTestRun(&teamHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "DELETE",
				queryURL:     fmt.Sprintf("/api/v1/team/%v", testCase.queryParamId),
				isBodyExists: false,
			})
		})
	}
}
