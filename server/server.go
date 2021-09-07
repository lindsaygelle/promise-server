package server

import (
	"database/sql"
	"log"
	"net/http"
)

func readRequest(request *http.Request) {
	log.Println(request.URL, request.Header)
}

func statusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if err == sql.ErrNoRows {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
