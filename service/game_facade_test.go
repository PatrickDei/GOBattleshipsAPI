package service

import (
	"github.com/golang/mock/gomock"
	"go-battleships/domain"
	"go-battleships/errors"
	"go-battleships/mocks/service"
	"testing"
)

var mockGameService *service.MockGameService
var mockPlayerService *service.MockPlayerService
var gf GameFacade

func gameFacadeSetup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockGameService = service.NewMockGameService(ctrl)
	mockPlayerService = service.NewMockPlayerService(ctrl)

	gf = GameFacadeImpl{gameService: mockGameService, playerService: mockPlayerService}

	return func() {
		defer ctrl.Finish()
	}
}

func TestChallengeOpponentChecksIfPlayersExist(t *testing.T) {
	teardown := gameFacadeSetup(t)
	defer teardown()

	playerId := "1"
	opponentId := "2"

	mockPlayerService.EXPECT().ExistsById(playerId).Return(true, nil).Times(1)
	mockPlayerService.EXPECT().ExistsById(opponentId).Return(false, nil).Times(1)

	gf.ChallengeOpponent(playerId, opponentId)
}

func TestChallengeOpponentReturnsNotFound(t *testing.T) {
	teardown := gameFacadeSetup(t)
	defer teardown()

	playerId := "1"
	opponentId := "2"

	mockPlayerService.EXPECT().ExistsById(playerId).Return(true, nil)
	mockPlayerService.EXPECT().ExistsById(opponentId).Return(false, errors.NewNotFoundError(errors.NewErrorBody("code", "arg")))

	if _, err := gf.ChallengeOpponent(playerId, opponentId); err == nil {
		t.Error("Not found was thrown but not returned")
	}
}

func TestChallengeOpponentReturnsGame(t *testing.T) {
	teardown := gameFacadeSetup(t)
	defer teardown()

	playerId := "1"
	opponentId := "2"
	g := domain.Game{
		Id:         "1",
		PlayerId:   playerId,
		OpponentId: opponentId,
		TurnCount:  0,
	}

	mockPlayerService.EXPECT().ExistsById(playerId).Return(true, nil)
	mockPlayerService.EXPECT().ExistsById(opponentId).Return(true, nil)
	mockGameService.EXPECT().CreateGame(playerId, opponentId).Return(&g, nil)

	if game, err := gf.ChallengeOpponent(playerId, opponentId); err != nil || game == nil {
		t.Error("Service returned game but facade didn't")
	}
}

func TestChallengeOpponentReturnsRuntimeError(t *testing.T) {
	teardown := gameFacadeSetup(t)
	defer teardown()

	playerId := "1"
	opponentId := "2"

	mockPlayerService.EXPECT().ExistsById(playerId).Return(true, nil)
	mockPlayerService.EXPECT().ExistsById(opponentId).Return(true, nil)
	mockGameService.EXPECT().CreateGame(playerId, opponentId).Return(nil, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	if _, err := gf.ChallengeOpponent(playerId, opponentId); err == nil {
		t.Error("Service returned error but facade didn't")
	}

}
