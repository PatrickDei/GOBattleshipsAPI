package domain

import (
	"database/sql"
	"github.com/PatrickDei/log-lib/logger"
	"github.com/jmoiron/sqlx"
	"go-battleships/errors"
	"strconv"
)

type GameRepositoryImpl struct {
	dbClient *sqlx.DB
}

func (gr GameRepositoryImpl) Save(g Game) (*Game, *errors.AppError) {
	insertStatement := "INSERT INTO Games (PlayerId, OpponentId, TurnCount, PlayerBoardId, OpponentBoardId, Status) VALUES (?, ?, ?, ?, ?, ?)"

	result, err := gr.dbClient.Exec(insertStatement, g.PlayerId, g.OpponentId, g.TurnCount, g.PlayerBoardId, g.OpponentBoardId, g.Status)
	if err != nil {
		logger.Error("Error while creating game")
		return nil, errors.NewInternalServerError(errors.NewErrorBody("error.db", "Error while creating game"))
	}

	id := extractId(result)

	g.Id = strconv.FormatInt(id, 10)

	return &g, nil
}

func (gr GameRepositoryImpl) GetById(id string) (*Game, *errors.AppError) {
	selectStatement := "SELECT PlayerId, OpponentId, TurnCount, Status FROM Games WHERE Id = ?"

	var g Game
	err := gr.dbClient.Get(&g, selectStatement, id)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err.Error())
			return nil, errors.NewInternalServerError(errors.NewErrorBody("error.db", "Error while reading game"))
		} else {
			return nil, errors.NewNotFoundError(errors.NewErrorBody("unknown-game-id", id))
		}
	}

	return &g, nil
}

func NewGameRepository(dbClient *sqlx.DB) GameRepository {
	return GameRepositoryImpl{dbClient: dbClient}
}
