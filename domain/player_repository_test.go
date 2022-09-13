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

func TestExistsByIdReturnsResult(t *testing.T) {
	mock, teardown := setup()
	defer teardown()

	id := "1"

	mock.ExpectQuery("SELECT 1 FROM Players WHERE").WillReturnRows(
		sqlmock.NewRows([]string{"id"}))

	if exists, _ := pr.ExistsById(id); exists != false {
		t.Error("Database returned no rows but repo returned true")
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

/*
func TestGetByIdReturnsPlayer(t *testing.T) {
	mock, teardown := setup()
	defer teardown()

	id := "1"
	p := Player{
		Id:    id,
		Name:  "John",
		Email: "doe@mail.com",
	}

	mock.ExpectQuery("SELECT Name, Email FROM Players WHERE").WillReturnRows(
		sqlmock.NewRows([]string{"Name", "Email"}).AddRow("john", "mail@mail.com"))

	pa, e := pr.GetById(id)
	fmt.Printf("%+v %+v", pa, e)
	if player, err := pr.GetById(id); player.Id != p.Id || err != nil {
		t.Error("The method didn't return the requested Id")
	}
}*/

func TestGetAllReturnsPlayers(t *testing.T) {
	mock, teardown := setup()
	defer teardown()

	mock.ExpectQuery("SELECT Name, Email FROM Players").WillReturnRows(
		sqlmock.NewRows([]string{"Name", "Email"}).AddRow("John", "Doe"))

	if _, err := pr.GetAll(); err != nil {
		t.Error("Database returned rows but repo returned error")
	}
}

func TestGetAllThrowsError(t *testing.T) {
	mock, teardown := setup()
	defer teardown()

	mock.ExpectQuery("SELECT Name, Email FROM Players").WillReturnError(sql.ErrTxDone)

	if _, err := pr.GetAll(); err == nil {
		t.Error("Database returned error but repo didn't")
	}
}
