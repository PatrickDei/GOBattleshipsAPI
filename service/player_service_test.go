package service

import (
	"github.com/golang/mock/gomock"
	"go-battleships/errors"
	"go-battleships/mocks/domain"
	"testing"
)

var mockRepo *domain.MockPlayerRepository
var ps PlayerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockPlayerRepository(ctrl)

	ps = PlayerServiceImpl{repo: mockRepo}

	return func() {
		defer ctrl.Finish()
	}
}

func TestExistsByEmailReturnsRepositoryResponse(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	email := "pperic@mail.com"

	mockRepo.EXPECT().ExistsByEmail(email).Return(true, nil)

	if exists, _ := ps.ExistsByEmail(email); exists != true {
		t.Error("Repository returned true but service returned false")
	}

	mockRepo.EXPECT().ExistsByEmail(email).Return(false, nil)

	if exists, _ := ps.ExistsByEmail(email); exists != false {
		t.Error("Repository returned false but service returned true")
	}
}

func TestExistsByEmailReturnsRepositoryError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	email := "pperic@mail.com"

	mockRepo.EXPECT().ExistsByEmail(email).Return(false, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	if _, err := ps.ExistsByEmail(email); err == nil {
		t.Error("Repository returned error but service didn't pass it")
	}
}
