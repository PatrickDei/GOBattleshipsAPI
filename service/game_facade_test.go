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
var mockBoardService *service.MockBoardService
var gf GameFacade

func gameFacadeSetup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockGameService = service.NewMockGameService(ctrl)
	mockPlayerService = service.NewMockPlayerService(ctrl)
	mockBoardService = service.NewMockBoardService(ctrl)

	gf = GameFacadeImpl{gameService: mockGameService, playerService: mockPlayerService, boardService: mockBoardService}

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
	boardId := "1"
	g := domain.Game{
		Id:         "1",
		PlayerId:   playerId,
		OpponentId: opponentId,
		TurnCount:  0,
	}
	b := domain.NewEmptyBoard()
	b.Id = boardId

	mockPlayerService.EXPECT().ExistsById(playerId).Return(true, nil)
	mockPlayerService.EXPECT().ExistsById(opponentId).Return(true, nil)
	mockBoardService.EXPECT().CreateNewBoard().Return(&b, nil).Times(2)
	mockGameService.EXPECT().CreateGame(playerId, opponentId, boardId, boardId).Return(&g, nil)

	if game, err := gf.ChallengeOpponent(playerId, opponentId); err != nil || game == nil {
		t.Error("Service returned game but facade didn't")
	}
}

func TestChallengeOpponentReturnsRuntimeError(t *testing.T) {
	teardown := gameFacadeSetup(t)
	defer teardown()

	playerId := "1"
	opponentId := "2"
	boardId := "1"
	b := domain.NewEmptyBoard()
	b.Id = boardId

	mockPlayerService.EXPECT().ExistsById(playerId).Return(true, nil)
	mockPlayerService.EXPECT().ExistsById(opponentId).Return(true, nil)
	mockBoardService.EXPECT().CreateNewBoard().Return(&b, nil).Times(2)
	mockGameService.EXPECT().CreateGame(playerId, opponentId, boardId, boardId).Return(nil, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	if _, err := gf.ChallengeOpponent(playerId, opponentId); err == nil {
		t.Error("Service returned error but facade didn't")
	}

}

func TestGetGameStatusReturnsErrorWhenPlayerNotFound(t *testing.T) {
	teardown := gameFacadeSetup(t)
	defer teardown()

	mockPlayerService.EXPECT().ExistsById(gomock.Any()).Return(false, nil)

	if _, err := gf.GetGameStatus("", ""); err == nil {
		t.Error("Player was not found but facade didn't return error")
	}
}

func TestGetGameStatusReturnsErrorWhenGameNotFound(t *testing.T) {
	teardown := gameFacadeSetup(t)
	defer teardown()

	mockPlayerService.EXPECT().ExistsById(gomock.Any()).Return(true, nil)
	mockGameService.EXPECT().GetById(gomock.Any()).Return(nil, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	if _, err := gf.GetGameStatus("", ""); err == nil {
		t.Error("Game was not found but facade didn't return error")
	}
}

func TestGetGameStatusReturnsErrorWhenPlayerBoardNotFound(t *testing.T) {
	teardown := gameFacadeSetup(t)
	defer teardown()

	g := domain.Game{}

	mockPlayerService.EXPECT().ExistsById(gomock.Any()).Return(true, nil)
	mockGameService.EXPECT().GetById(gomock.Any()).Return(&g, nil)
	mockBoardService.EXPECT().GetByPlayerIdAndGameId(gomock.Any(), gomock.Any()).Return(nil, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	if _, err := gf.GetGameStatus("", ""); err == nil {
		t.Error("Board was not found but facade didn't return error")
	}
}

func TestGetGameStatusReturnsGameState(t *testing.T) {
	teardown := gameFacadeSetup(t)
	defer teardown()

	g := domain.Game{}
	b := domain.Board{}

	mockPlayerService.EXPECT().ExistsById(gomock.Any()).Return(true, nil)
	mockGameService.EXPECT().GetById(gomock.Any()).Return(&g, nil)
	mockBoardService.EXPECT().GetByPlayerIdAndGameId(gomock.Any(), gomock.Any()).Return(&b, nil)

	if gs, _ := gf.GetGameStatus("", ""); gs == nil {
		t.Error("Facade didn't return any game state")
	}
}

func TestListPlayersByGameReturnsError(t *testing.T) {
	teardown := gameFacadeSetup(t)
	defer teardown()

	mockGameService.EXPECT().ListByPlayerId(gomock.Any()).Return(nil, errors.NewInternalServerError(errors.NewErrorBody("code", "err")))

	if _, err := gf.ListPlayersGames(""); err == nil {
		t.Error("Service returned error but facade didn't")
	}
}

func TestListPlayersByGameReturnsListOfGames(t *testing.T) {
	teardown := gameFacadeSetup(t)
	defer teardown()

	g := []domain.Game{
		{}, {},
	}

	mockGameService.EXPECT().ListByPlayerId(gomock.Any()).Return(g, nil)

	if games, _ := gf.ListPlayersGames(""); len(games) == 0 {
		t.Error("Service returned games but facade didn't")
	}
}
