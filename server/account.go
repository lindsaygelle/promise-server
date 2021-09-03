package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/account"
	"github.com/lindsaygelle/promise/promise-server/database"
)

func RouteAccount(client database.Client, engine *gin.Engine) {
	engine.GET("/account", func(c *gin.Context) {
		accounts, err := account.GetAccounts(client)
		c.JSON(statusCode(err), accounts)
	})
	engine.GET("/account/:id", func(c *gin.Context) {
		account, err := account.GetAccount(client, c.Param("id"))
		c.JSON(statusCode(err), account)
	})
	engine.GET("/account/:id/setting", func(c *gin.Context) {
		setting, err := account.GetSetting(client, c.Param("id"))
		c.JSON(statusCode(err), setting)
	})
}
