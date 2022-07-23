// Code generated by MockGen. DO NOT EDIT.
// Source: internal/team/usecase.go

// Package mock_team is a generated GoMock package.
package mock_team

import (
	reflect "reflect"

	models "github.com/Unlites/nba_api/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUseCase) Create(stat *models.Team) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", stat)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUseCaseMockRecorder) Create(stat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), stat)
}

// Delete mocks base method.
func (m *MockUseCase) Delete(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUseCaseMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), id)
}

// GetById mocks base method.
func (m *MockUseCase) GetById(id int64) (*models.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUseCaseMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUseCase)(nil).GetById), id)
}

// Update mocks base method.
func (m *MockUseCase) Update(team *models.Team) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", team)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUseCaseMockRecorder) Update(team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUseCase)(nil).Update), team)
}
