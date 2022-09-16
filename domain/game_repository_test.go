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

	mock.ExpectExec("^INSERT INTO Games").WillReturnResult(sqlmock.NewResult(lastInsertId, 1))

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

	mock.ExpectExec("^INSERT INTO Games").WillReturnError(sql.ErrTxDone)

	if _, err := gr.Save(g); err == nil {
		t.Error("An error occurred when querying the database but it wasn't thrown")
	}
}

func TestListByPlayerIdReturnsSliceOfGames(t *testing.T) {
	mock, teardown := gameSetup()
	defer teardown()

	g := Game{
		Id:         "1",
		PlayerId:   "1",
		OpponentId: "2",
		Status:     InProgress,
	}

	mock.ExpectQuery("^SELECT Id, PlayerId, OpponentId, Status FROM Games WHERE").WillReturnRows(
		sqlmock.NewRows(
			[]string{"Id", "PlayerId", "OpponentId", "Status"}).AddRow(g.Id, g.PlayerId, g.OpponentId, g.Status))

	if g, _ := gr.ListByPlayerId(""); len(g) == 0 {
		t.Error("Db returned rows but repository returned an empty slice")
	}
}

func TestListByPlayerIdReturnsError(t *testing.T) {
	mock, teardown := gameSetup()
	defer teardown()

	mock.ExpectQuery("^SELECT Id, PlayerId, OpponentId, Status FROM Games WHERE").WillReturnError(sql.ErrTxDone)

	if _, err := gr.ListByPlayerId(""); err == nil {
		t.Error("Db returned error but repository didn't")
	}
}
