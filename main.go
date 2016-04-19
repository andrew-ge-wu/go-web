package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"fmt"
	"github.com/andrew-ge-wu/go-web/storage"
	"gopkg.in/mgo.v2"
)

var dataStore storage.DataStore

func main() {

	glb_session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer glb_session.Close()

	// Optional. Switch the session to a monotonic behavior.
	glb_session.SetMode(mgo.Monotonic, true)
	dataStore = storage.DataStore{glb_session}
	router := gin.Default()
	router.GET("/v1/country/:country/eco/:ecoId/socialdata/:source", read)
	router.Run() // listen and server on 0.0.0.0:8080
}

func read(c *gin.Context) {
	country := c.Param("country")
	ecoId := strings.Split(c.Param("ecoId"), ",")
	source := strings.Split(c.Param("source"), ",")
	action := c.Query("action")
	//query := c.Query("query")
	fmt.Print(action)
	result := dataStore.GET(country, ecoId, source)
	if &result != nil {
		c.JSON(http.StatusOK, result)
	} else {
		c.Status(http.StatusNotFound)
	}
}
