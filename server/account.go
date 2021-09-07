package server

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/account"
)

func RouteAccount(database *sql.DB, engine *gin.Engine) {
	engine.GET("/account", func(c *gin.Context) {
		accounts, err := account.ReadProfiles(database)
		c.JSON(statusCode(err), accounts)
		log.Println(err)
		readRequest(c.Request)
	})
	engine.GET("/account/:id", func(c *gin.Context) {
		account, err := account.ReadProfile(database, c.Param("id"))
		c.JSON(statusCode(err), account)
		log.Println(err)
		readRequest(c.Request)
	})
}
