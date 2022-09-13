package service

import (
	"go-battleships/domain"
	"go-battleships/errors"
)

type GameService interface {
	CreateGame(playerId string, opponentId string) (*domain.Game, *errors.AppError)
}

type GameServiceImpl struct {
	repo domain.GameRepository
}

func (gs GameServiceImpl) CreateGame(playerId string, opponentId string) (*domain.Game, *errors.AppError) {
	g := domain.Game{
		PlayerId:   playerId,
		OpponentId: opponentId,
		TurnCount:  0,
	}

	savedGame, err := gs.repo.Save(g)
	if err != nil {
		return nil, err
	}

	return savedGame, nil
}

func NewGameService(repo domain.GameRepository) GameService {
	return GameServiceImpl{repo: repo}
}
