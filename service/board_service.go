package service

import (
	"go-battleships/domain"
	"go-battleships/errors"
)

//go:generate mockgen -destination=../mocks/service/mock_board_service.go -package=service -source=board_service.go BoardService
type BoardService interface {
	CreateNewBoardForPlayer(string) (*domain.Board, *errors.AppError)
	GetByPlayerIdAndGameId(playerId string, gameId string) (*domain.Board, *errors.AppError)
}

type BoardServiceImpl struct {
	repo         domain.BoardRepository
	boardFactory domain.BoardFactory
}

func (bs BoardServiceImpl) CreateNewBoardForPlayer(playerId string) (*domain.Board, *errors.AppError) {
	b := bs.boardFactory.GenerateNewBoard()
	b.PlayerId = playerId

	board, err := bs.repo.Save(b)
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (bs BoardServiceImpl) GetByPlayerIdAndGameId(playerId string, gameId string) (*domain.Board, *errors.AppError) {
	return bs.repo.GetByPlayerIdAndGameId(playerId, gameId)
}

func NewBoardService(br domain.BoardRepository, bf domain.BoardFactory) BoardService {
	return BoardServiceImpl{repo: br, boardFactory: bf}
}
