package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/location"
)

func RouteLocation(client database.Client, engine *gin.Engine) {
	engine.GET("/location/country", func(c *gin.Context) {
		countries, err := location.GetCountries(client)
		c.JSON(statusCode(err), countries)
	})
	engine.GET("/location/country/:id", func(c *gin.Context) {
		country, err := location.GetCountry(client, c.Param("id"))
		c.JSON(statusCode(err), country)
	})
}
