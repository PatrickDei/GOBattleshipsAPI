package domain

import (
	"go-battleships/errors"
	"strings"
)

const BoardColumnCount = 10
const BoardRowCount = 10

type Board struct {
	Id     string
	Fields string
}

func (b Board) GetFieldsAsSlice() []string {
	rows := make([]string, 0)
	row := ""
	for i, r := range b.Fields {
		if i%BoardColumnCount == 0 {
			rows = append(rows, row)
			row = ""
		} else {
			row = row + string(r)
		}
	}
	return rows
}

func NewEmptyBoard() Board {
	return Board{Fields: strings.Repeat(strings.Repeat(".", BoardColumnCount), BoardRowCount)}
}

//go:generate mockgen -destination=../mocks/domain/mock_board_repository.go -package=domain -source=board.go BoardRepository
type BoardRepository interface {
	Save(Board) (*Board, *errors.AppError)
}
