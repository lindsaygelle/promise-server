package server

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/account"
)

func RouteAccount(database *sql.DB, engine *gin.Engine) {
	engine.GET("/account", func(c *gin.Context) {
		readRequest(c.Request)
		accounts, err := account.ReadProfiles(database)
		c.JSON(statusCode(err), accounts)
		log.Println(err)
	})
	engine.POST("/account", func(c *gin.Context) {
		readRequest(c.Request)
		account, err := account.WriteProfile(database, c.Param("id"), c.Request.Body)
		c.JSON(statusCode(err), account)
		log.Println(err)
	})
	engine.GET("/account/:id", func(c *gin.Context) {
		readRequest(c.Request)
		account, err := account.ReadProfile(database, c.Param("id"))
		c.JSON(statusCode(err), account)
		log.Println(err)
	})
}
