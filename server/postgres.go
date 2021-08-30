package server

import (
	"net/http"

	"github.com/lindsaygelle/promise/promise-server/postgres"
)

func Postgres(w http.ResponseWriter, r *http.Request, v postgres.Client) {
	v.Ping()
}
