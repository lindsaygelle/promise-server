package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
		accounts, err := server.GetAccounts(postgres)
		statusCode := http.StatusOK
		log.Println(c.Request.URL.Path, err)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, accounts)
	})
	r.GET("/account/:id", func(c *gin.Context) {
		id := c.Param("id")
		account, err := server.GetAccount(postgres, id)
		statusCode := http.StatusOK
		log.Println(c.Request.URL.Path, id, err)
		if err != nil {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, account)
	})
	r.GET("/account/:id/setting", func(c *gin.Context) {
		id := c.Param("id")
		setting, err := server.GetAccountSetting(postgres, id)
		statusCode := http.StatusOK
		log.Println(c.Request.URL.Path, id, err)
		if err != nil {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, setting)
	})
	r.GET("/account/setting", func(c *gin.Context) {
		settings, err := server.GetAccountSettings(postgres)
		statusCode := http.StatusOK
		log.Println(c.Request.URL.Path, err)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, settings)
	})
	r.GET("/language", func(c *gin.Context) {
		languages, err := server.GetLanguages(postgres)
		statusCode := http.StatusOK
		log.Println(c.Request.URL.Path, err)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, languages)
	})
	r.GET("/language/:id", func(c *gin.Context) {
		id := c.Param("id")
		language, err := server.GetLanguage(postgres, id)
		statusCode := http.StatusOK
		log.Println(c.Request.URL.Path, id, err)
		if err != nil {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, language)
	})
	r.GET("/language/:id/tag", func(c *gin.Context) {
		id := c.Param("id")
		tag, err := server.GetLanguageTag(postgres, id)
		statusCode := http.StatusOK
		log.Println(c.Request.URL.Path, id, err)
		if err != nil {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, tag)
	})
	r.GET("/location/country", func(c *gin.Context) {
		countries, err := server.GetCountries(postgres)
		statusCode := http.StatusOK
		log.Println(c.Request.URL.Path, err)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, countries)
	})
	r.GET("/location/country/:id", func(c *gin.Context) {
		id := c.Param("id")
		country, err := server.GetCountry(postgres, id)
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
		log.Println(c.Request.URL.Path, err)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, nil)
	})
	r.GET("/redis", func(c *gin.Context) {
		_, err := redis.Ping().Result()
		statusCode := http.StatusOK
		log.Println(c.Request.URL.Path, err)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, nil)
	})
	defer postgres.Close()
	defer redis.Close()
	r.Run(fmt.Sprintf(":%s", os.Getenv("ADDR")))
}
