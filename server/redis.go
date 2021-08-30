package server

import (
	"net/http"

	"github.com/lindsaygelle/promise/promise-server/redis"
)

func Redis(w http.ResponseWriter, r *http.Request, v redis.Client) {
	v.Ping()
}
