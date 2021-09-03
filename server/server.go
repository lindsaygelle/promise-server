package server

import (
	"database/sql"
	"net/http"
)

func statusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if err == sql.ErrNoRows {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
