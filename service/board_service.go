package service

import (
	"go-battleships/domain"
	"go-battleships/errors"
)

//go:generate mockgen -destination=../mocks/service/mock_board_service.go -package=service -source=board_service.go BoardService
type BoardService interface {
	CreateNewBoard() (*domain.Board, *errors.AppError)
}

type BoardServiceImpl struct {
	repo         domain.BoardRepository
	boardFactory domain.BoardFactory
}

func (bs BoardServiceImpl) CreateNewBoard() (*domain.Board, *errors.AppError) {
	b := domain.NewEmptyBoard()
	bs.boardFactory.PopulateBoard(&b)

	board, err := bs.repo.Save(b)
	if err != nil {
		return nil, err
	}
	return board, nil
}

func NewBoardService(br domain.BoardRepository, bf domain.BoardFactory) BoardService {
	return BoardServiceImpl{repo: br, boardFactory: bf}
}
