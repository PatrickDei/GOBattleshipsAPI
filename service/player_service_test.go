package service

import (
	"github.com/golang/mock/gomock"
	realdomain "go-battleships/domain"
	realdto "go-battleships/dto"
	"go-battleships/errors"
	"go-battleships/mocks/domain"
	"testing"
)

var mockPlayerRepo *domain.MockPlayerRepository
var ps PlayerService
var pc = realdto.PlayerCommand{
	Name:  "name",
	Email: "email",
}

func playerSetup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockPlayerRepo = domain.NewMockPlayerRepository(ctrl)

	ps = PlayerServiceImpl{repo: mockPlayerRepo}

	return func() {
		defer ctrl.Finish()
	}
}

func TestGetByIdReturnsPlayer(t *testing.T) {
	teardown := playerSetup(t)
	defer teardown()

	id := "1"
	player := realdomain.Player{
		Id:    "1",
		Name:  "John",
		Email: "Doe@mail.com",
	}

	mockPlayerRepo.EXPECT().GetById(id).Return(&player, nil)

	p, err := ps.GetById(id)
	if err != nil {
		t.Error("Repo returned player but service returned error")
	}
	if p.Id != "" {
		t.Error("Service delivered Id from the db")
	}
	if p.Name == "" || p.Email == "" {
		t.Error("Not all data delivered from the database was delivered")
	}
}

func TestGetByIdReturnsError(t *testing.T) {
	teardown := playerSetup(t)
	defer teardown()

	id := "1"

	mockPlayerRepo.EXPECT().GetById(id).Return(nil, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	p, err := ps.GetById(id)
	if err == nil {
		t.Error("Repo returned error but service didn't")
	}
	if p != nil {
		t.Error("Repo returned nil for player but service didn't")
	}
}

func TestExistsByEmailReturnsRepositoryResponse(t *testing.T) {
	teardown := playerSetup(t)
	defer teardown()

	email := "pperic@mail.com"

	mockPlayerRepo.EXPECT().ExistsByEmail(email).Return(true, nil)

	if exists, _ := ps.ExistsByEmail(email); exists != true {
		t.Error("Repository returned true but service returned false")
	}

	mockPlayerRepo.EXPECT().ExistsByEmail(email).Return(false, nil)

	if exists, _ := ps.ExistsByEmail(email); exists != false {
		t.Error("Repository returned false but service returned true")
	}
}

func TestExistsByEmailReturnsRepositoryError(t *testing.T) {
	teardown := playerSetup(t)
	defer teardown()

	email := "pperic@mail.com"

	mockPlayerRepo.EXPECT().ExistsByEmail(email).Return(false, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	if _, err := ps.ExistsByEmail(email); err == nil {
		t.Error("Repository returned error but service didn't pass it")
	}
}

func TestCreatePlayerPassesInternalServerError(t *testing.T) {
	teardown := playerSetup(t)
	defer teardown()

	mockPlayerRepo.EXPECT().ExistsByEmail(pc.Email).Return(false, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	if _, err := ps.CreatePlayer(pc); err == nil {
		t.Error("Repository returned error but service didn't pass it")
	}

	mockPlayerRepo.EXPECT().ExistsByEmail(pc.Email).Return(false, nil)
	mockPlayerRepo.EXPECT().Save(gomock.Any()).Return(nil, errors.NewInternalServerError(errors.NewErrorBody("code", "arg")))

	if _, err := ps.CreatePlayer(pc); err == nil {
		t.Error("Repository returned error but service didn't pass it")
	}
}

func TestCreatePlayerThrowsConflictWhenPlayerExists(t *testing.T) {
	teardown := playerSetup(t)
	defer teardown()

	mockPlayerRepo.EXPECT().ExistsByEmail(pc.Email).Return(true, nil)

	if _, err := ps.CreatePlayer(pc); err == nil {
		t.Error("Player exists but no conflict was thrown")
	}
}

func TestCreatePlayerReturnsDTOAfterSave(t *testing.T) {
	teardown := playerSetup(t)
	defer teardown()

	p := realdomain.Player{}.FromCommand(pc)

	mockPlayerRepo.EXPECT().ExistsByEmail(pc.Email).Return(false, nil)
	mockPlayerRepo.EXPECT().Save(p).Return(&p, nil)

	if dto, err := ps.CreatePlayer(pc); dto == nil || *dto != p.ToDTO() || err != nil {
		t.Error("Player created but an error was returned")
	}
}

func TestGetAllReturnsPlayers(t *testing.T) {
	teardown := playerSetup(t)
	defer teardown()

	mockPlayerRepo.EXPECT().GetAll().Return([]realdomain.Player{
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
	teardown := playerSetup(t)
	defer teardown()

	errBody := errors.ErrorBody{"Error": "DB error"}

	mockPlayerRepo.EXPECT().GetAll().Return(nil, errors.NewInternalServerError(errBody))

	if player, err := ps.GetAll(); player != nil || err == nil {
		t.Error("Repository returned players but service encountered an error")
	}
}
