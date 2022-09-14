package service

import (
	"github.com/golang/mock/gomock"
	realdomain "go-battleships/domain"
	"go-battleships/errors"
	"go-battleships/mocks/domain"
	"testing"
)

var mockGameRepo *domain.MockGameRepository
var gs GameService

func gameSetup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockGameRepo = domain.NewMockGameRepository(ctrl)

	gs = GameServiceImpl{repo: mockGameRepo}

	return func() {
		defer ctrl.Finish()
	}
}

func TestCreateGameReturnsSavedGame(t *testing.T) {
	teardown := gameSetup(t)
	defer teardown()

	playerId := "1"
	opponentId := "2"
	g := realdomain.Game{
		PlayerId:   playerId,
		OpponentId: opponentId,
		TurnCount:  0,
	}
	gWithId := g
	gWithId.Id = "1"

	mockGameRepo.EXPECT().Save(g).Return(&gWithId, nil)

	savedG, err := gs.CreateGame(playerId, opponentId, "", "")
	if err != nil {
		t.Error("Repository returned game but service returned error")
	}
	if savedG == nil {
		t.Error("Repository returned game but service returned nil")
	}
}

func TestCreateGameReturnsError(t *testing.T) {
	teardown := gameSetup(t)
	defer teardown()

	playerId := "1"
	opponentId := "2"
	g := realdomain.Game{
		PlayerId:   playerId,
		OpponentId: opponentId,
		TurnCount:  0,
	}

	mockGameRepo.EXPECT().Save(g).Return(nil, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	savedG, err := gs.CreateGame(playerId, opponentId, "", "")
	if err == nil {
		t.Error("Repository returned error but service didn't")
	}
	if savedG != nil {
		t.Error("Repository returned error but service returned game")
	}
}
