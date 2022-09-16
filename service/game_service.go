package service

import (
	"go-battleships/domain"
	"go-battleships/errors"
)

//go:generate mockgen -destination=../mocks/service/mock_game_service.go -package=service -source=game_service.go GameService
type GameService interface {
	CreateGame(playerId string, opponentId string, playerBoardId string, opponentBoardId string) (*domain.Game, *errors.AppError)
	GetById(string) (*domain.Game, *errors.AppError)
	ListByPlayerId(string) ([]domain.Game, *errors.AppError)
}

type GameServiceImpl struct {
	repo domain.GameRepository
}

func (gs GameServiceImpl) CreateGame(playerId string, opponentId string, playerBoardId string, opponentBoardId string) (*domain.Game, *errors.AppError) {
	g := domain.Game{
		PlayerId:        playerId,
		OpponentId:      opponentId,
		TurnCount:       0,
		PlayerBoardId:   playerBoardId,
		OpponentBoardId: opponentBoardId,
		Status:          domain.InProgress,
	}

	savedGame, err := gs.repo.Save(g)
	if err != nil {
		return nil, err
	}

	return savedGame, nil
}

func (gs GameServiceImpl) GetById(id string) (*domain.Game, *errors.AppError) {
	return gs.repo.GetById(id)
}

func (gs GameServiceImpl) ListByPlayerId(playerId string) ([]domain.Game, *errors.AppError) {
	g, err := gs.repo.ListByPlayerId(playerId)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func NewGameService(repo domain.GameRepository) GameService {
	return GameServiceImpl{repo: repo}
}
