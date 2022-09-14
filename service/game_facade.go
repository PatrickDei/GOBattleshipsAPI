package service

import (
	"go-battleships/dto"
	"go-battleships/errors"
)

//go:generate mockgen -destination=../mocks/service/mock_game_facade.go -package=service -source=game_facade.go GameFacade
type GameFacade interface {
	ChallengeOpponent(playerId string, opponentId string) (*dto.GameDTO, *errors.AppError)
}

type GameFacadeImpl struct {
	gameService   GameService
	playerService PlayerService
	boardService  BoardService
}

func (gf GameFacadeImpl) ChallengeOpponent(playerId string, opponentId string) (*dto.GameDTO, *errors.AppError) {
	if exists, err := gf.playerExistsById(playerId); exists != true {
		return nil, err
	}

	if exists, err := gf.playerExistsById(opponentId); exists != true {
		return nil, err
	}

	playerBoard, err := gf.boardService.CreateNewBoard()
	if err != nil {
		return nil, err
	}

	opponentBoard, err := gf.boardService.CreateNewBoard()
	if err != nil {
		return nil, err
	}

	g, err := gf.gameService.CreateGame(playerId, opponentId, playerBoard.Id, opponentBoard.Id)
	if err != nil {
		return nil, err
	}

	resp := g.ToDTO()

	return &resp, nil
}

func (gf GameFacadeImpl) playerExistsById(id string) (bool, *errors.AppError) {
	if exists, err := gf.playerService.ExistsById(id); exists != true || err != nil {
		if err != nil {
			return false, err
		}
		return false, errors.NewNotFoundError(errors.NewErrorBody("unknown-user-id", id))
	}
	return true, nil
}

func NewGameFacade(s GameService, p PlayerService, b BoardService) GameFacade {
	return GameFacadeImpl{gameService: s, playerService: p, boardService: b}
}
