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

func (pr PlayerRepositoryImpl) ExistsByEmail(email string) (bool, *errors.AppError) {
	selectStatement := "SELECT 1 FROM Players WHERE Email = ?"

	var exists bool
	err := pr.dbClient.QueryRow(selectStatement, email).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while checking if player with email exists: " + err.Error())
		return false, errors.NewInternalServerError(map[string]string{"Message": "User not found"})
	}

	return exists, nil
}

func extractId(r sql.Result) int64 {
	id, err := r.LastInsertId()
	if err != nil {
		logger.Error("Error while extracting new id")
	}

	return id
}

func NewPlayerRepository(dbClient *sqlx.DB) PlayerRepository {
	return PlayerRepositoryImpl{dbClient: dbClient}
}
