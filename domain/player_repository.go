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
	}

	id := extractId(result)

	p.Id = strconv.FormatInt(id, 10)

	return &p, nil
}

func (pr PlayerRepositoryImpl) FindByEmail(email string) ([]Player, *errors.AppError) {
	selectStatement := "SELECT * FROM Players WHERE Email = ?"

	p := make([]Player, 0)
	err := pr.dbClient.Select(&p, selectStatement, email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError(map[string]string{"Message": "User not found"})
		} else {
			logger.Error("Error while scanning user: " + err.Error())
			return nil, errors.NewInternalServerError(map[string]string{"Message": "User not found"})
		}
	}

	return p, nil
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
