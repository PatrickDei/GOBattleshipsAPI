package service

import (
	"github.com/golang/mock/gomock"
	realdomain "go-battleships/domain"
	realdto "go-battleships/dto"
	"go-battleships/errors"
	"go-battleships/mocks/domain"
	"testing"
)

var mockRepo *domain.MockPlayerRepository
var ps PlayerService
var pc = realdto.PlayerCommand{
	Name:  "name",
	Email: "email",
}

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

func TestCreatePlayerPassesInternalServerError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockRepo.EXPECT().ExistsByEmail(pc.Email).Return(false, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	if _, err := ps.CreatePlayer(pc); err == nil {
		t.Error("Repository returned error but service didn't pass it")
	}

	mockRepo.EXPECT().ExistsByEmail(pc.Email).Return(false, nil)
	mockRepo.EXPECT().Save(gomock.Any()).Return(nil, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	if _, err := ps.CreatePlayer(pc); err == nil {
		t.Error("Repository returned error but service didn't pass it")
	}
}

func TestCreatePlayerThrowsConflictWhenPlayerExists(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockRepo.EXPECT().ExistsByEmail(pc.Email).Return(true, nil)

	if _, err := ps.CreatePlayer(pc); err == nil {
		t.Error("Player exists but no conflict was thrown")
	}
}

func TestCreatePlayerReturnsDTOAfterSave(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	p := realdomain.Player{}.FromCommand(pc)

	mockRepo.EXPECT().ExistsByEmail(pc.Email).Return(false, nil)
	mockRepo.EXPECT().Save(p).Return(&p, nil)

	if dto, err := ps.CreatePlayer(pc); dto == nil || *dto != p.ToDTO() || err != nil {
		t.Error("Player created but an error was returned")
	}
}

func TestGetAllReturnsPlayers(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockRepo.EXPECT().GetAll().Return([]realdto.PlayerDTO{
		{
			Name:  "john",
			Email: "doe@mail.com",
		},
	}, nil)

	if player, err := ps.GetAll(); player == nil || err != nil {
		t.Error("Repository returned players but service encountered an error")
	}
}

func TestGetAllReturnsThrownError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	errBody := errors.ErrorBody{"Error": "DB error"}

	mockRepo.EXPECT().GetAll().Return(nil, errors.NewInternalServerError(errBody))

	if player, err := ps.GetAll(); player != nil || err == nil {
		t.Error("Repository returned players but service encountered an error")
	}
}
