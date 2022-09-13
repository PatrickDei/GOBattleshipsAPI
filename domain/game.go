package domain

import (
	"go-battleships/dto"
	"go-battleships/errors"
)

type Game struct {
	Id         string
	PlayerId   string
	OpponentId string
	TurnCount  int
}

func (g Game) ToDTO() dto.GameDTO {
	return dto.GameDTO{
		Id:         g.Id,
		PlayerId:   g.PlayerId,
		OpponentId: g.OpponentId,
		Starting:   g.determineStartingId(),
	}
}

func (g Game) determineStartingId() string {
	if g.TurnCount%2 == 0 {
		return g.PlayerId
	}
	return g.OpponentId
}

//go:generate mockgen -destination=../mocks/domain/mock_game_repository.go -package=domain -source=game.go GameRepository
type GameRepository interface {
	Save(Game) (*Game, *errors.AppError)
}
