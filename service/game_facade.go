package service

import (
	"go-battleships/dto"
	"go-battleships/errors"
)

type GameFacade interface {
	ChallengeOpponent(playerId string, opponentId string) (*dto.GameDTO, *errors.AppError)
}

type GameFacadeImpl struct {
	gameService GameService
}

func (gf GameFacadeImpl) ChallengeOpponent(playerId string, opponentId string) (*dto.GameDTO, *errors.AppError) {
	// checks for existence

	g, err := gf.gameService.CreateGame(playerId, opponentId)
	if err != nil {
		return nil, err
	}

	resp := g.ToDTO()

	return &resp, nil
}

func NewGameFacade(s GameService) GameFacade {
	return GameFacadeImpl{gameService: s}
}
