package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
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
	return &redis.Options{
		Addr:     (os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0}
}

func main() {
	redisClientOptions := getRedisOptions()
	redisClient := getRedis(redisClientOptions)
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
