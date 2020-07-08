package main

import (
	"github.com/gin-gonic/gin"
	"go_mongodb/api/modules"
	"net/http"
)

func main() {

	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/url.insert", Insert)
	}
	router.Run(":8080")
}

func Insert(c *gin.Context) {

	song := modules.Song{
		Title:     "Track01",
		Artist:    "NoName",
		Album:     "Unknown",
		Performer: "Unknown",
	}

	resultInsert := modules.InsertDB(song)
	c.JSON(http.StatusOK, resultInsert)
}
