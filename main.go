package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"fmt"
)

func main() {
	router := gin.Default()
	router.GET("/v1/country/:country/eco/:ecoId/socialdata/:source", func(c *gin.Context) {
		country := c.Param("country")
		ecoId := strings.Split(c.Param("ecoId"), ",")
		source := strings.Split(c.Param("source"), ",")
		action := c.Query("action")
		query := c.Query("query")
		fmt.Print(action)
		c.String(http.StatusOK, read(country, ecoId, source, query))
	})
	router.Run() // listen and server on 0.0.0.0:8080
}

func read(country string,
ecoId []string,
source []string,
query string) (string) {
	return country + "|" + strings.Join(ecoId, "/") + "|" + strings.Join(source, "/") + "|" + query
}
