package server

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/account"
)

func RouteAccount(database *sql.DB, engine *gin.Engine) {
	engine.GET("/account", func(c *gin.Context) {
		accounts, err := account.ReadAccounts(database)
		log.Println(err)
		c.JSON(statusCode(err), accounts)
	})
	engine.GET("/account/:id", func(c *gin.Context) {
		account, err := account.ReadAccount(database, c.Param("id"))
		log.Println(err)
		c.JSON(statusCode(err), account)
	})
	engine.GET("/account/:id/setting", func(c *gin.Context) {
		setting, err := account.ReadSetting(database, c.Param("id"))
		log.Println(err)
		c.JSON(statusCode(err), setting)
	})
	engine.POST("/account", func(c *gin.Context) {
		account, err := account.WriteAccount(database, c.Request)
		log.Println(err)
		c.JSON(statusCode(err), account)
	})
	engine.PUT("/account/:id/name", func(c *gin.Context) {
		err := account.EditAccountName(database, c.Param("id"), c.Request)
		log.Println(err)
		c.JSON(statusCode(err), nil)
	})
}
