package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

func getRedis(redisOptions *redis.Options) *redis.Client {
	redisClient := redis.NewClient(redisOptions)
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	return redisClient
}

func getRedisOptions() *redis.Options {
	var (
		host     = os.Getenv("REDIS_HOST")
		password = os.Getenv("REDIS_PASSWORD")
		port     = os.Getenv("REDIS_PORT")
	)
	return &redis.Options{
		Addr:     (host + ":" + port),
		Password: password,
		DB:       0}
}

func getPostgres() *sql.DB {
	var (
		database = os.Getenv("POSTGRES_DB")
		host     = os.Getenv("POSTGRES_HOST")
		password = os.Getenv("POSTGRES_PASSWORD")
		port     = os.Getenv("POSTGRES_PORT")
		sslmode  = os.Getenv("POSTGRES_SSL")
		user     = os.Getenv("POSTGRES_USER")
	)
	dataSourceName := fmt.Sprintf("dbname=%s host=%s password=%s port=%s sslmode=%s user=%s", database, host, password, port, sslmode, user)
	driverName := "postgres"
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	redisClientOptions := getRedisOptions()
	redisClient := getRedis(redisClientOptions)
	postgresClient := getPostgres()
	http.HandleFunc("/postgres", func(w http.ResponseWriter, r *http.Request) {
		err := postgresClient.Ping()
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
		s, err := redisClient.Ping().Result()
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
