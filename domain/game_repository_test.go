package domain

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"strconv"
	"testing"
)

var gr GameRepository

func gameSetup() (sqlmock.Sqlmock, func()) {
	mockDb, mock, _ := sqlmock.New()

	gr = GameRepositoryImpl{dbClient: sqlx.NewDb(mockDb, "mysql")}

	return mock, func() {
		defer mockDb.Close()
	}
}

func TestSaveReturnsSavedGame(t *testing.T) {
	mock, teardown := gameSetup()
	defer teardown()

	lastInsertId := int64(10001)
	g := Game{
		PlayerId:   "1",
		OpponentId: "2",
		TurnCount:  0,
	}

	mock.ExpectExec("INSERT INTO Games").WillReturnResult(sqlmock.NewResult(lastInsertId, 1))

	if savedGame, _ := gr.Save(g); savedGame == nil || savedGame.Id != strconv.FormatInt(lastInsertId, 10) {
		t.Error("The returned game doesn't contain the last saved Id to db")
	}
}

func TestSaveGameReturnsRuntimeError(t *testing.T) {
	mock, teardown := gameSetup()
	defer teardown()

	g := Game{
		PlayerId:   "1",
		OpponentId: "2",
		TurnCount:  0,
	}

	mock.ExpectExec("INSERT INTO Games").WillReturnError(sql.ErrTxDone)

	if _, err := gr.Save(g); err == nil {
		t.Error("An error occurred when querying the database but it wasn't thrown")
	}
}
