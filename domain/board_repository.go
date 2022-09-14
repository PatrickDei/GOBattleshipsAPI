package domain

import (
	"github.com/PatrickDei/log-lib/logger"
	"github.com/jmoiron/sqlx"
	"go-battleships/errors"
	"strconv"
)

type BoardRepositoryImpl struct {
	dbClient *sqlx.DB
}

func (br BoardRepositoryImpl) Save(b Board) (*Board, *errors.AppError) {
	insertStatement := "INSERT INTO Boards (Fields) VALUES (?)"

	result, err := br.dbClient.Exec(insertStatement, b.Fields)
	if err != nil {
		logger.Error("Error while creating board")
		return nil, errors.NewInternalServerError(errors.NewErrorBody("error.db", "Error while creating board"))
	}

	id := extractId(result)

	b.Id = strconv.FormatInt(id, 10)

	return &b, nil
}

func NewBoardRepository(db *sqlx.DB) BoardRepository {
	return BoardRepositoryImpl{dbClient: db}
}
