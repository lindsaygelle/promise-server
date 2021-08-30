package main

import (
	_ "github.com/lib/pq"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lindsaygelle/promise/promise-server/postgres"
	"github.com/lindsaygelle/promise/promise-server/redis"
)

func main() {
	redis := redis.NewClient(redis.NewConfig())
	postgres := postgres.NewClient(postgres.NewConfig())
	http.HandleFunc("/postgres", func(w http.ResponseWriter, r *http.Request) {
		err := postgres.Ping()
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(struct {
			StatusCode int
		}{
			StatusCode: statusCode,
		})
	})
	http.HandleFunc("/redis", func(w http.ResponseWriter, r *http.Request) {
		s, err := redis.Ping().Result()
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(struct {
			Content    string
			StatusCode int
		}{
			Content:    s,
			StatusCode: statusCode,
		})
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			StatusCode int
		}{
			StatusCode: http.StatusOK,
		})
	})
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("ADDR")), nil))
}
