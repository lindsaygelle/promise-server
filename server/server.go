package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/lindsaygelle/w3g"
)

func readRequest(request *http.Request) {
	request.Header.Add(w3g.XRequestID, uuid.New().String())
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
