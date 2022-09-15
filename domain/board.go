package domain

import (
	"go-battleships/errors"
	"strings"
)

const BoardColumnCount = 10
const BoardRowCount = 10

type Board struct {
	Id        string `db:"Id"`
	Fields    string `db:"Fields"`
	ShipCount int    `db:"ShipCount"`
}

func (b Board) GetFieldsAsSlice() []string {
	rows := make([]string, 0)
	row := ""
	for i, r := range b.Fields {
		row = row + string(r)
		if i%BoardColumnCount == 9 {
			rows = append(rows, row)
			row = ""
		}
	}
	return rows
}

func NewEmptyBoard() Board {
	return Board{Fields: strings.Repeat(strings.Repeat(Unpopulated.String(), BoardColumnCount), BoardRowCount)}
}

//go:generate mockgen -destination=../mocks/domain/mock_board_repository.go -package=domain -source=board.go BoardRepository
type BoardRepository interface {
	Save(Board) (*Board, *errors.AppError)
	GetByPlayerIdAndGameId(playerId string, gameId string) (*Board, *errors.AppError)
}
