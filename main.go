package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
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
