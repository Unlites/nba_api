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
	mock_stat "github.com/Unlites/nba_api/internal/stat/mocks"
	httpErr "github.com/Unlites/nba_api/pkg/http_errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name                 string
	queryParamId         int64
	queryBody            *models.Stat
	mockBehavior         func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat)
	expectedStatusCode   int
	expectedResponseBody string
}

type statHandlerTestParams struct {
	t            *testing.T
	testCase     testCase
	queryMethod  string
	queryURL     string
	isBodyExists bool
}

func statHandlerTestRun(testParams *statHandlerTestParams) {
	ctrl := gomock.NewController(testParams.t)
	defer ctrl.Finish()

	uc := mock_stat.NewMockUseCase(ctrl)
	testParams.testCase.mockBehavior(uc, testParams.testCase.queryParamId, testParams.testCase.queryBody)

	handler := NewStatHandler(uc)

	router := gin.Default()
	group := router.Group("/api/v1/stat")

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
			mockBehavior: func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {
				s.EXPECT().GetById(queryParamId).Return(&models.Stat{
					Id:       1,
					GameId:   1,
					PlayerId: 1,
					Points:   "30",
					Rebounds: "5",
					Assists:  "5",
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"game_id":1,"player_id":1,"points":"30","rebounds":"5","assists":"5"}`,
		},
		{
			name:                 "Not positive id",
			queryParamId:         0,
			mockBehavior:         func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrNotPositiveId,
		},
		{
			name:         "Not found",
			queryParamId: 2,
			mockBehavior: func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {
				s.EXPECT().GetById(queryParamId).Return(nil, errors.New("sql: no rows in result set"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: httpErr.ErrNotFound,
		},
		{
			name:         "Internal error",
			queryParamId: 1,
			mockBehavior: func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {
				s.EXPECT().GetById(queryParamId).Return(nil, errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			statHandlerTestRun(&statHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "GET",
				queryURL:     fmt.Sprintf("/api/v1/stat/%v", testCase.queryParamId),
				isBodyExists: false,
			})
		})
	}
}

func TestHandler_Create(t *testing.T) {
	testTable := []testCase{
		{
			name: "Success",
			queryBody: &models.Stat{
				GameId:   1,
				PlayerId: 1,
				Points:   "30",
				Rebounds: "5",
				Assists:  "5",
			},
			mockBehavior: func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {
				s.EXPECT().Create(queryBody).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name: "Missed param",
			queryBody: &models.Stat{
				GameId:   1,
				Points:   "30",
				Rebounds: "5",
				Assists:  "5",
			},
			mockBehavior:         func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name: "Empty param",
			queryBody: &models.Stat{
				GameId:   1,
				PlayerId: 0,
				Points:   "30",
				Rebounds: "5",
				Assists:  "5",
			},
			mockBehavior:         func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name: "Internal error",
			queryBody: &models.Stat{
				GameId:   1,
				PlayerId: 1,
				Points:   "30",
				Rebounds: "5",
				Assists:  "5",
			},
			mockBehavior: func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {
				s.EXPECT().Create(queryBody).Return(errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
		{
			name: "Not positive param",
			queryBody: &models.Stat{
				GameId:   1,
				PlayerId: 1,
				Points:   "-30",
				Rebounds: "5",
				Assists:  "5",
			},
			mockBehavior:         func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			statHandlerTestRun(&statHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "POST",
				queryURL:     "/api/v1/stat/",
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
			queryBody: &models.Stat{
				GameId:   1,
				PlayerId: 1,
				Points:   "30",
				Rebounds: "5",
				Assists:  "5",
			},
			mockBehavior: func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {
				s.EXPECT().Update(queryBody).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:         "Not positive id",
			queryParamId: 0,
			queryBody: &models.Stat{
				GameId:   1,
				PlayerId: 1,
				Points:   "30",
				Rebounds: "5",
				Assists:  "5",
			},
			mockBehavior:         func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrNotPositiveId,
		},
		{
			name:         "Not found",
			queryParamId: 1,
			queryBody: &models.Stat{
				GameId:   1,
				PlayerId: 1,
				Points:   "30",
				Rebounds: "5",
				Assists:  "5",
			},
			mockBehavior: func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {
				s.EXPECT().Update(queryBody).Return(errors.New("sql: no rows in result set"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: httpErr.ErrNotFound,
		},
		{
			name:         "Missed param",
			queryParamId: 1,
			queryBody: &models.Stat{
				GameId:   1,
				Points:   "30",
				Rebounds: "5",
				Assists:  "5",
			},
			mockBehavior:         func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name:         "Empty param",
			queryParamId: 1,
			queryBody: &models.Stat{
				GameId:   1,
				PlayerId: 0,
				Points:   "30",
				Rebounds: "5",
				Assists:  "5",
			},
			mockBehavior:         func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrInvalidJSON,
		},
		{
			name:         "Internal error",
			queryParamId: 1,
			queryBody: &models.Stat{
				GameId:   1,
				PlayerId: 1,
				Points:   "30",
				Rebounds: "5",
				Assists:  "5",
			},
			mockBehavior: func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {
				s.EXPECT().Update(queryBody).Return(errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			statHandlerTestRun(&statHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "PUT",
				queryURL:     fmt.Sprintf("/api/v1/stat/%v", testCase.queryParamId),
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
			mockBehavior: func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {
				s.EXPECT().Delete(queryParamId).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:                 "Not positive id",
			queryParamId:         0,
			mockBehavior:         func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {},
			expectedStatusCode:   400,
			expectedResponseBody: httpErr.ErrNotPositiveId,
		},
		{
			name:         "Not found",
			queryParamId: 2,
			mockBehavior: func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {
				s.EXPECT().Delete(queryParamId).Return(errors.New("sql: no rows in result set"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: httpErr.ErrNotFound,
		},
		{
			name:         "Internal error",
			queryParamId: 1,
			mockBehavior: func(s *mock_stat.MockUseCase, queryParamId int64, queryBody *models.Stat) {
				s.EXPECT().Delete(queryParamId).Return(errors.New("some error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: httpErr.ErrInternal,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			statHandlerTestRun(&statHandlerTestParams{
				t:            t,
				testCase:     testCase,
				queryMethod:  "DELETE",
				queryURL:     fmt.Sprintf("/api/v1/stat/%v", testCase.queryParamId),
				isBodyExists: false,
			})
		})
	}
}
