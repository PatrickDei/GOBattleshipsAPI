package service

import (
	"go-battleships/domain"
	"go-battleships/dto"
	"go-battleships/errors"
)

const UsernameTaken = "username-already-taken"

//go:generate mockgen -destination=../mocks/service/mock_player_service.go -package=service -source=player_service.go PlayerService
type PlayerService interface {
	CreatePlayer(command dto.PlayerCommand) (*dto.PlayerDTO, *errors.AppError)
	ExistsByEmail(string) (bool, *errors.AppError)
	GetById(string) (*dto.PlayerDTO, *errors.AppError)
}

type PlayerServiceImpl struct {
	repo domain.PlayerRepository
}

func (ps PlayerServiceImpl) CreatePlayer(pc dto.PlayerCommand) (*dto.PlayerDTO, *errors.AppError) {
	p := domain.Player{}.FromCommand(pc)

	playerExists, err := ps.ExistsByEmail(pc.Email)
	if err != nil {
		return nil, err
	}
	if playerExists {
		return nil, errors.NewConflictError(errors.NewErrorBody(UsernameTaken, pc.Email))
	}

	player, err := ps.repo.Save(p)
	if err != nil {
		return nil, err
	}

	resp := player.ToDTO()

	return &resp, nil
}

func (ps PlayerServiceImpl) ExistsByEmail(email string) (bool, *errors.AppError) {
	exists, err := ps.repo.ExistsByEmail(email)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (ps PlayerServiceImpl) GetById(id string) (*dto.PlayerDTO, *errors.AppError) {
	p, err := ps.repo.GetById(id)
	pdto := p.ToDTO()
	pdto.Id = ""
	return &pdto, err
}

func NewPlayerService(repo domain.PlayerRepository) PlayerService {
	return PlayerServiceImpl{repo: repo}
}
