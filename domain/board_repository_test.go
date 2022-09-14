package domain

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
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
	b := Board{
		fields: strings.Repeat(strings.Repeat(".", BoardColumnCount), BoardRowCount),
	}

	mock.ExpectExec("INSERT INTO Boards").WillReturnResult(sqlmock.NewResult(lastInsertId, 1))

	if board, _ := br.Save(b); board == nil || board.Id != strconv.FormatInt(lastInsertId, 10) {
		t.Error("The returned board doesn't contain the last saved Id to db")
	}
}

func TestSaveBoardReturnsRuntimeError(t *testing.T) {
	mock, teardown := boardSetup()
	defer teardown()

	b := Board{
		fields: strings.Repeat(strings.Repeat(".", BoardColumnCount), BoardRowCount),
	}

	mock.ExpectExec("INSERT INTO Boards").WillReturnError(sql.ErrTxDone)

	if _, err := br.Save(b); err == nil {
		t.Error("An error occurred when querying the database but it wasn't thrown")
	}
}
