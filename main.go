package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lindsaygelle/promise/promise-server/account"
	"github.com/lindsaygelle/promise/promise-server/location"
	"github.com/lindsaygelle/promise/promise-server/postgres"
	"github.com/lindsaygelle/promise/promise-server/redis"
	"github.com/lindsaygelle/promise/promise-server/server"
)

func main() {
	var (
		postgres = postgres.NewClient(postgres.NewConfig())
		redis    = redis.NewClient(redis.NewConfig())
	)
	r := gin.Default()
	r.GET("/account", func(c *gin.Context) {
		accounts, err := server.Accounts(postgres)
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, accounts)
	})
	r.GET("/account/:id", func(c *gin.Context) {
		id := c.Param("id")
		account, err := account.GetAccount(postgres, id)
		log.Println(c.Request.URL.Path, id, err)
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, account)
	})
	r.GET("/account/:id/setting", func(c *gin.Context) {
		id := c.Param("id")
		setting, err := account.GetSetting(postgres, id)
		log.Println(c.Request.URL.Path, id, err)
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, setting)
	})
	r.GET("/location/country", func(c *gin.Context) {
		countries, err := location.GetCountries(postgres)
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, countries)
	})
	r.GET("/location/country/:id", func(c *gin.Context) {
		id := c.Param("id")
		country, err := location.GetCountry(postgres, id)
		statusCode := http.StatusOK
		log.Println(c.Request.URL.Path, id, err)
		if err != nil {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, country)
	})
	r.GET("/postgres", func(c *gin.Context) {
		err := postgres.Ping()
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, nil)
	})
	r.GET("/redis", func(c *gin.Context) {
		_, err := redis.Ping().Result()
		statusCode := http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, nil)
	})
	defer postgres.Close()
	defer redis.Close()
	r.Run(fmt.Sprintf(":%s", os.Getenv("ADDR")))
}
