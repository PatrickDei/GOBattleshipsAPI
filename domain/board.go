package domain

import "go-battleships/errors"

const BoardColumnCount = 10
const BoardRowCount = 10

type Board struct {
	Id     string
	fields string
}

func (b Board) GetFields() []string {
	rows := make([]string, 0)
	row := ""
	for i, r := range b.fields {
		if i%BoardColumnCount == 0 {
			rows = append(rows, row)
			row = ""
		} else {
			row = row + string(r)
		}
	}
	return rows
}

//go:generate mockgen -destination=../mocks/domain/mock_board_repository.go -package=domain -source=board.go BoardRepository
type BoardRepository interface {
	Save(Board) (*Board, *errors.AppError)
}
