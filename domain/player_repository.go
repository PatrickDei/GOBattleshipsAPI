package domain

import (
	"database/sql"
	"github.com/PatrickDei/log-lib/logger"
	"github.com/jmoiron/sqlx"
	"go-battleships/errors"
	"strconv"
)

type PlayerRepositoryImpl struct {
	dbClient *sqlx.DB
}

func (pr PlayerRepositoryImpl) Save(p Player) (*Player, *errors.AppError) {
	insertStatement := "INSERT INTO Players (Name, Email) VALUES (?, ?)"

	result, err := pr.dbClient.Exec(insertStatement, p.Name, p.Email)
	if err != nil {
		logger.Error("Error while creating player")
		return nil, errors.NewInternalServerError(map[string]string{"error": "An unexpected error occurred"})
	}

	id := extractId(result)

	p.Id = strconv.FormatInt(id, 10)

	return &p, nil
}

func extractId(r sql.Result) int64 {
	id, err := r.LastInsertId()
	if err != nil {
		logger.Error("Error while extracting new id")
	}

	return id
}

func (pr PlayerRepositoryImpl) ExistsByEmail(email string) (bool, *errors.AppError) {
	return pr.existsByColumn("Email", email)
}

func (pr PlayerRepositoryImpl) ExistsById(id string) (bool, *errors.AppError) {
	return pr.existsByColumn("Id", id)
}

func (pr PlayerRepositoryImpl) existsByColumn(column string, value string) (bool, *errors.AppError) {
	selectStatement := "SELECT 1 FROM Players WHERE " + column + " = ?"

	var exists bool
	err := pr.dbClient.QueryRow(selectStatement, value).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while checking if player with " + column + " exists: " + err.Error())
		return false, errors.NewInternalServerError(map[string]string{"Message": "User not found"})
	}

	return exists, nil
}

func (pr PlayerRepositoryImpl) GetById(id string) (*Player, *errors.AppError) {
	selectStatement := "SELECT * FROM Players WHERE Id = ?"

	var p Player
	err := pr.dbClient.Get(&p, selectStatement, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError(nil)
		} else {
			logger.Error("Error while reading player")
			return nil, errors.NewInternalServerError(map[string]string{"error": "An unexpected error occurred"})
		}
	}
	return &p, nil
}

func (pr PlayerRepositoryImpl) GetAll() ([]Player, *errors.AppError) {
	selectStatement := "SELECT Name, Email FROM Players"

	p := make([]Player, 0)

	err := pr.dbClient.Select(&p, selectStatement)
	if err != nil {
		logger.Error("Error while fetching all players: " + err.Error())
		return nil, errors.NewInternalServerError(map[string]string{"Message": "Encountered error when getting available players"})
	}

	return p, nil
}

func NewPlayerRepository(dbClient *sqlx.DB) PlayerRepository {
	return PlayerRepositoryImpl{dbClient: dbClient}
}
