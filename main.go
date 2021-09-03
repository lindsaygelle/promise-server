package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/postgres"
	"github.com/lindsaygelle/promise/promise-server/redis"
	"github.com/lindsaygelle/promise/promise-server/server"
)

func main() {
	var (
		engine   = gin.Default()
		postgres = postgres.NewClient(postgres.NewConfig())
		redis    = redis.NewClient(redis.NewConfig())
	)
	for _, fn := range []func(database.Client, *gin.Engine){
		server.RouteAccount,
		server.RouteLanguage,
		server.RoutePromise} {
		fn(postgres, engine)
	}

	defer postgres.Close()
	defer redis.Close()

	engine.Run(fmt.Sprintf(":%s", os.Getenv("ADDR")))
}
