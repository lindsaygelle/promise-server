package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/language"
)

func RouteLanguage(client database.Client, engine *gin.Engine) {
	engine.GET("/language", func(c *gin.Context) {
		languages, err := language.GetLanguages(client)
		c.JSON(statusCode(err), languages)
	})
	engine.GET("/language/:id", func(c *gin.Context) {
		language, err := language.GetLanguage(client, c.Param("id"))
		c.JSON(statusCode(err), language)
	})
}
