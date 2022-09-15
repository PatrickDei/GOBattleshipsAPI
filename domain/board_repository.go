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
	insertStatement := "INSERT INTO Boards (Fields, ShipCount, PlayerId) VALUES (?, ?, ?)"

	result, err := br.dbClient.Exec(insertStatement, b.Fields, b.ShipCount, b.PlayerId)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.NewInternalServerError(errors.NewErrorBody("error.db", "Error while creating board"))
	}

	id := extractId(result)

	b.Id = strconv.FormatInt(id, 10)

	return &b, nil
}

func (br BoardRepositoryImpl) GetByPlayerIdAndGameId(playerId string, gameId string) (*Board, *errors.AppError) {
	selectStatement := "SELECT b.Id, b.Fields FROM Boards b JOIN Games G on b.Id = G.PlayerBoardId OR b.Id = G.OpponentBoardId WHERE G.Id = ? AND b.PlayerId = ?"

	var b Board
	err := br.dbClient.Get(&b, selectStatement, gameId, playerId)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.NewInternalServerError(errors.NewErrorBody("error.db", "Error while reading board"))
	}

	return &b, nil
}

func NewBoardRepository(db *sqlx.DB) BoardRepository {
	return BoardRepositoryImpl{dbClient: db}
}
