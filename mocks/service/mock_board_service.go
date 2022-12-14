// Code generated by MockGen. DO NOT EDIT.
// Source: board_service.go

// Package service is a generated GoMock package.
package service

import (
	domain "go-battleships/domain"
	errors "go-battleships/errors"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBoardService is a mock of BoardService interface.
type MockBoardService struct {
	ctrl     *gomock.Controller
	recorder *MockBoardServiceMockRecorder
}

// MockBoardServiceMockRecorder is the mock recorder for MockBoardService.
type MockBoardServiceMockRecorder struct {
	mock *MockBoardService
}

// NewMockBoardService creates a new mock instance.
func NewMockBoardService(ctrl *gomock.Controller) *MockBoardService {
	mock := &MockBoardService{ctrl: ctrl}
	mock.recorder = &MockBoardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBoardService) EXPECT() *MockBoardServiceMockRecorder {
	return m.recorder
}

// CreateNewBoard mocks base method.
func (m *MockBoardService) CreateNewBoardForPlayer(string) (*domain.Board, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewBoardForPlayer")
	ret0, _ := ret[0].(*domain.Board)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// CreateNewBoard indicates an expected call of CreateNewBoard.
func (mr *MockBoardServiceMockRecorder) CreateNewBoard() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewBoardForPlayer", reflect.TypeOf((*MockBoardService)(nil).CreateNewBoardForPlayer))
}

// GetByPlayerIdAndGameId mocks base method.
func (m *MockBoardService) GetByPlayerIdAndGameId(playerId, gameId string) (*domain.Board, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPlayerIdAndGameId", playerId, gameId)
	ret0, _ := ret[0].(*domain.Board)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// GetByPlayerIdAndGameId indicates an expected call of GetByPlayerIdAndGameId.
func (mr *MockBoardServiceMockRecorder) GetByPlayerIdAndGameId(playerId, gameId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPlayerIdAndGameId", reflect.TypeOf((*MockBoardService)(nil).GetByPlayerIdAndGameId), playerId, gameId)
}
