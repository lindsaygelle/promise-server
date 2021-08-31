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
	"github.com/lindsaygelle/promise/promise-server/server"
	"github.com/lindsaygelle/w3g"
)

func main() {
	redis := redis.NewClient(redis.NewConfig())
	postgres := postgres.NewClient(postgres.NewConfig())
	defer postgres.Close()
	http.HandleFunc("/postgres", func(w http.ResponseWriter, r *http.Request) {
		err := postgres.Ping()
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		w.Header().Set(w3g.ContentType, "application/json")
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
		w.Header().Set(w3g.ContentType, "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(struct {
			Content    string
			StatusCode int
		}{
			Content:    s,
			StatusCode: statusCode,
		})
	})
	http.HandleFunc("/account/accounts", func(w http.ResponseWriter, r *http.Request) {
		accounts, err := server.Accounts(postgres)
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		log.Println(r.URL.Path, err)
		w.Header().Set(w3g.ContentType, "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(accounts)
	})
	http.HandleFunc("/account/settings", func(w http.ResponseWriter, r *http.Request) {
		settings, err := server.AccountSettings(postgres)
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		log.Println(r.URL.Path, err)
		w.Header().Set(w3g.ContentType, "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(settings)
	})
	http.HandleFunc("/language/languages", func(w http.ResponseWriter, r *http.Request) {
		languages, err := server.Languages(postgres)
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		log.Println(r.URL.Path, err)
		w.Header().Set(w3g.ContentType, "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(languages)
	})
	http.HandleFunc("/language/tags", func(w http.ResponseWriter, r *http.Request) {
		tags, err := server.LanguageTags(postgres)
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		log.Println(r.URL.Path, err)
		w.Header().Set(w3g.ContentType, "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(tags)
	})
	http.HandleFunc("/location/countries", func(w http.ResponseWriter, r *http.Request) {
		countries, err := server.Countries(postgres)
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		log.Println(r.URL.Path, err)
		w.Header().Set(w3g.ContentType, "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(countries)
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
