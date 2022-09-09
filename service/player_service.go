package service

import (
	"go-battleships/domain"
	"go-battleships/dto"
	"go-battleships/errors"
)

type PlayerService interface {
	CreatePlayer(command dto.PlayerCommand) (*dto.PlayerDTO, *errors.AppError)
	FindByEmail(string) ([]dto.PlayerDTO, *errors.AppError)
}

type PlayerServiceImpl struct {
	repo domain.PlayerRepository
}

func (ps PlayerServiceImpl) CreatePlayer(pc dto.PlayerCommand) (*dto.PlayerDTO, *errors.AppError) {
	p := domain.Player{}.FromCommand(pc)

	existingPlayers, err := ps.FindByEmail(pc.Email)
	if existingPlayers != nil && len(existingPlayers) != 0 {
		return nil, errors.NewConflictError(map[string]string{
			"error-code": "error.username-already-taken",
			"error-arg":  "pero.peric@ag04.com",
		})
	}

	player, err := ps.repo.Save(p)
	if err != nil {
		return nil, err
	}

	resp := player.ToDTO()

	return &resp, nil
}

func (ps PlayerServiceImpl) FindByEmail(email string) ([]dto.PlayerDTO, *errors.AppError) {
	players, err := ps.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	resp := make([]dto.PlayerDTO, 0)
	for _, p := range players {
		resp = append(resp, p.ToDTO())
	}

	return resp, nil
}

func NewPlayerService(repo domain.PlayerRepository) PlayerService {
	return PlayerServiceImpl{repo: repo}
}
