package domain

import (
	"github.com/PatrickDei/log-lib/logger"
	"github.com/jmoiron/sqlx"
	"go-battleships/errors"
	"strconv"
)

type GameRepositoryImpl struct {
	dbClient *sqlx.DB
}

func (gr GameRepositoryImpl) Save(g Game) (*Game, *errors.AppError) {
	insertStatement := "INSERT INTO Games (PlayerId, OpponentId, TurnCount) VALUES (?, ?, ?)"

	result, err := gr.dbClient.Exec(insertStatement, g.PlayerId, g.OpponentId, g.TurnCount)
	if err != nil {
		logger.Error("Error while creating game")
		return nil, errors.NewInternalServerError(errors.NewErrorBody("error.db", "Error while creating game"))
	}

	id := extractId(result)

	g.Id = strconv.FormatInt(id, 10)

	return &g, nil
}

func NewGameRepository(dbClient *sqlx.DB) GameRepository {
	return GameRepositoryImpl{dbClient: dbClient}
}
