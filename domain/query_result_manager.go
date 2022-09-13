package domain

import (
	"database/sql"
	"github.com/PatrickDei/log-lib/logger"
)

func extractId(r sql.Result) int64 {
	id, err := r.LastInsertId()
	if err != nil {
		logger.Error("Error while extracting new id")
	}

	return id
}
