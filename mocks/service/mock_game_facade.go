// Code generated by MockGen. DO NOT EDIT.
// Source: game_facade.go

// Package service is a generated GoMock package.
package service

import (
	dto "go-battleships/dto"
	errors "go-battleships/errors"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGameFacade is a mock of GameFacade interface.
type MockGameFacade struct {
	ctrl     *gomock.Controller
	recorder *MockGameFacadeMockRecorder
}

// MockGameFacadeMockRecorder is the mock recorder for MockGameFacade.
type MockGameFacadeMockRecorder struct {
	mock *MockGameFacade
}

// NewMockGameFacade creates a new mock instance.
func NewMockGameFacade(ctrl *gomock.Controller) *MockGameFacade {
	mock := &MockGameFacade{ctrl: ctrl}
	mock.recorder = &MockGameFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameFacade) EXPECT() *MockGameFacadeMockRecorder {
	return m.recorder
}

// ChallengeOpponent mocks base method.
func (m *MockGameFacade) ChallengeOpponent(playerId, opponentId string) (*dto.GameDTO, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChallengeOpponent", playerId, opponentId)
	ret0, _ := ret[0].(*dto.GameDTO)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// ChallengeOpponent indicates an expected call of ChallengeOpponent.
func (mr *MockGameFacadeMockRecorder) ChallengeOpponent(playerId, opponentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChallengeOpponent", reflect.TypeOf((*MockGameFacade)(nil).ChallengeOpponent), playerId, opponentId)
}
