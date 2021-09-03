package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/promise"
)

func RoutePromise(client database.Client, engine *gin.Engine) {
	engine.GET("/promise", func(c *gin.Context) {
		promises, err := promise.GetPromises(client)
		c.JSON(statusCode(err), promises)
	})
	engine.GET("/promise/:id", func(c *gin.Context) {
		promise, err := promise.GetPromise(client, c.Param("id"))
		c.JSON(statusCode(err), promise)
	})
	engine.GET("/promise/:id/tag", func(c *gin.Context) {
		tags, err := promise.GetTagsPromise(client, c.Param("id"))
		c.JSON(statusCode(err), tags)
	})
	engine.GET("/promise/:id/vote", func(c *gin.Context) {
		votes, err := promise.GetVotes(client)
		c.JSON(statusCode(err), votes)
	})
	engine.GET("/promise/maker/:id", func(c *gin.Context) {
		promises, err := promise.GetPromisesMaker(client, c.Param("id"))
		c.JSON(statusCode(err), promises)
	})
	engine.GET("/promise/owner/:id", func(c *gin.Context) {
		promises, err := promise.GetPromisesOwner(client, c.Param("id"))
		c.JSON(statusCode(err), promises)
	})
}
