package domain

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"strconv"
	"testing"
)

var br BoardRepository

func boardSetup() (sqlmock.Sqlmock, func()) {
	mockDb, mock, _ := sqlmock.New()

	br = BoardRepositoryImpl{dbClient: sqlx.NewDb(mockDb, "mysql")}

	return mock, func() {
		defer mockDb.Close()
	}
}

func TestSaveReturnsSavedBoard(t *testing.T) {
	mock, teardown := boardSetup()
	defer teardown()

	lastInsertId := int64(10001)
	b := NewEmptyBoard()

	mock.ExpectExec("^INSERT INTO Boards").WillReturnResult(sqlmock.NewResult(lastInsertId, 1))

	if board, _ := br.Save(b); board == nil || board.Id != strconv.FormatInt(lastInsertId, 10) {
		t.Error("The returned board doesn't contain the last saved Id to db")
	}
}

func TestSaveBoardReturnsRuntimeError(t *testing.T) {
	mock, teardown := boardSetup()
	defer teardown()

	b := NewEmptyBoard()

	mock.ExpectExec("^INSERT INTO Boards").WillReturnError(sql.ErrTxDone)

	if _, err := br.Save(b); err == nil {
		t.Error("An error occurred when querying the database but it wasn't thrown")
	}
}

func TestGetByPlayerIdAndGameIdReturnsBoard(t *testing.T) {
	mock, teardown := boardSetup()
	defer teardown()

	b := NewEmptyBoard()
	b.Id = "1"

	mock.ExpectQuery("^SELECT (.*) FROM Boards . JOIN Games .").WillReturnRows(sqlmock.NewRows([]string{"Id", "Fields"}).AddRow(b.Id, b.Fields))

	if b, _ := br.GetByPlayerIdAndGameId("", ""); b == nil {
		t.Error("Db returned a Board but repository didn't return it")
	}
}

func TestGetByPlayerIdAndGameIdReturnsError(t *testing.T) {
	mock, teardown := boardSetup()
	defer teardown()

	mock.ExpectQuery("^SELECT (.*) FROM Boards . JOIN Games .").WillReturnError(sql.ErrTxDone)

	if _, err := br.GetByPlayerIdAndGameId("", ""); err == nil {
		t.Error("An error occurred when fetching the board but repository didn't return it")
	}
}
