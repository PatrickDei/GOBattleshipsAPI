package domain

import (
	"go-battleships/dto"
	"go-battleships/errors"
)

type Game struct {
	Id              string `db:"Id"`
	PlayerId        string `db:"PlayerId"`
	OpponentId      string `db:"OpponentId"`
	TurnCount       int    `db:"TurnCount"`
	PlayerBoardId   string
	OpponentBoardId string
	Status          GameStatus `db:"Status"`
}

func (g Game) ToDTO() dto.GameDTO {
	return dto.GameDTO{
		Id:         g.Id,
		PlayerId:   g.PlayerId,
		OpponentId: g.OpponentId,
		Starting:   g.DetermineIdOfPlayersTurn(),
	}
}

func (g Game) ToStateDTOForPlayer(playerId string) dto.GameDTO {
	return dto.GameDTO{
		Id:         g.Id,
		PlayerId:   g.PlayerId,
		OpponentId: g.OpponentId,
		Status:     g.DetermineStatusForPlayer(playerId),
	}
}

func NewGameStateDTO(playerId string, g Game, b Board) dto.GameStateDTO {
	opponentId := determineOpponentId(playerId, g)
	gs := determineGameState(g)
	return dto.GameStateDTO{
		Player: dto.BoardState{
			PlayerId: playerId,
			Board:    b.GetFieldsAsSlice(),
		},
		Opponent: dto.BoardState{
			PlayerId: opponentId,
			Board:    NewEmptyBoard().GetFieldsAsSlice(),
		},
		Game: gs,
	}
}

func determineOpponentId(playerId string, g Game) string {
	var opponentId string
	if g.PlayerId == playerId {
		opponentId = g.OpponentId
	} else {
		opponentId = g.PlayerId
	}
	return opponentId
}

func determineGameState(g Game) dto.GameState {
	var gs dto.GameState

	if g.Status == InProgress {
		gs = dto.GameState{
			PlayerTurn: g.DetermineIdOfPlayersTurn(),
		}
	} else if g.Status == Finished {
		gs = dto.GameState{
			Won: g.DetermineIdOfPlayersTurn(),
		}
	}

	return gs
}

func (g Game) DetermineIdOfPlayersTurn() string {
	if g.TurnCount%2 == 0 {
		return g.PlayerId
	}
	return g.OpponentId
}

func (g Game) DetermineStatusForPlayer(playerId string) string {
	if g.Status == InProgress {
		return NotFinished.String()
	} else {
		if g.DetermineIdOfPlayersTurn() == playerId {
			return Won.String()
		}
		return Lost.String()
	}
}

//go:generate mockgen -destination=../mocks/domain/mock_game_repository.go -package=domain -source=game.go GameRepository
type GameRepository interface {
	Save(Game) (*Game, *errors.AppError)
	GetById(string) (*Game, *errors.AppError)
	ListByPlayerId(string) ([]Game, *errors.AppError)
}
