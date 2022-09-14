package service

import (
	"github.com/golang/mock/gomock"
	realdomain "go-battleships/domain"
	"go-battleships/errors"
	"go-battleships/mocks/domain"
	"testing"
)

var mockBoardRepo *domain.MockBoardRepository
var bs BoardService

func boardSetup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockBoardRepo = domain.NewMockBoardRepository(ctrl)

	bs = BoardServiceImpl{repo: mockBoardRepo}

	return func() {
		defer ctrl.Finish()
	}
}

func TestCreateNewBoardReturnsCreatedBoard(t *testing.T) {
	teardown := boardSetup(t)
	defer teardown()

	b := realdomain.NewEmptyBoard()
	boardWithId := b
	boardWithId.Id = "1"

	mockBoardRepo.EXPECT().Save(b).Return(&boardWithId, nil)

	if board, err := bs.CreateNewBoard(); board == nil || err != nil {
		t.Error("Repo returned board but service didn't")
	}
}

func TestCreateNewBoardReturnsError(t *testing.T) {
	teardown := boardSetup(t)
	defer teardown()

	b := realdomain.NewEmptyBoard()

	mockBoardRepo.EXPECT().Save(b).Return(nil, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	if _, err := bs.CreateNewBoard(); err == nil {
		t.Error("Repo returned error but service didn't")
	}
}
