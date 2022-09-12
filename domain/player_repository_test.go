package domain

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"strconv"
	"testing"
)

var pr PlayerRepository

func setup() (sqlmock.Sqlmock, func()) {
	mockDb, mock, _ := sqlmock.New()

	pr = PlayerRepositoryImpl{dbClient: sqlx.NewDb(mockDb, "mysql")}

	return mock, func() {
		defer mockDb.Close()
	}
}

func TestExistsByEmailReturnsTrue(t *testing.T) {
	mock, teardown := setup()
	defer teardown()

	email := "email"

	mock.ExpectQuery("SELECT 1 FROM Players WHERE").WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(1))

	if exists, _ := pr.ExistsByEmail(email); exists != true {
		t.Error("Database returned rows but repo returned false")
	}
}

func TestExistsByEmailReturnsFalse(t *testing.T) {
	mock, teardown := setup()
	defer teardown()

	email := "email"

	mock.ExpectQuery("SELECT 1 FROM Players WHERE").WillReturnRows(
		sqlmock.NewRows([]string{"id"}))

	if exists, _ := pr.ExistsByEmail(email); exists != false {
		t.Error("Database returned no rows but repo returned true")
	}
}

func TestExistsByEmailReturnsRuntimeError(t *testing.T) {
	mock, teardown := setup()
	defer teardown()

	email := "email"

	mock.ExpectQuery("SELECT 1 FROM Players WHERE").WillReturnError(sql.ErrTxDone)

	if _, err := pr.ExistsByEmail(email); err == nil {
		t.Error("An error occurred when querying the database but it wasn't thrown")
	}
}

func TestSaveReturnsSavedPlayer(t *testing.T) {
	mock, teardown := setup()
	defer teardown()

	lastInsertId := int64(10001)
	p := Player{
		Name:  "John",
		Email: "doe@mail.com",
	}

	mock.ExpectExec("INSERT INTO Players").WillReturnResult(sqlmock.NewResult(lastInsertId, 1))

	if savedPlayer, _ := pr.Save(p); savedPlayer == nil || savedPlayer.Id != strconv.FormatInt(lastInsertId, 10) {
		t.Error("The returned player doesn't contain the last saved Id to db")
	}
}

func TestSaveReturnsRuntimeError(t *testing.T) {
	mock, teardown := setup()
	defer teardown()

	p := Player{
		Name:  "John",
		Email: "doe@mail.com",
	}

	mock.ExpectExec("INSERT INTO Players").WillReturnError(sql.ErrTxDone)

	if _, err := pr.Save(p); err == nil {
		t.Error("An error occurred when querying the database but it wasn't thrown")
	}
}
