package domain

import (
	"go-battleships/dto"
	"go-battleships/errors"
)

type Player struct {
	Id    string `db:"Id"`
	Name  string `db:"Name"`
	Email string `db:"Email"`
}

func (p Player) FromCommand(c dto.PlayerCommand) Player {
	return Player{
		Name:  c.Name,
		Email: c.Email,
	}
}

func (p Player) ToDTO() dto.PlayerDTO {
	return dto.PlayerDTO{
		Id:    p.Id,
		Name:  p.Name,
		Email: p.Email,
	}
}

type PlayerRepository interface {
	Save(Player) (*Player, *errors.AppError)
	ExistsByEmail(string) (bool, *errors.AppError)
}
