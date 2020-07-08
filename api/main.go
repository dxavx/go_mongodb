package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go_mongodb/api/modules"
	"log"
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

	client, ctx, _ := modules.ConnectMongo()
	defer client.Disconnect(ctx)

	// List DB
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	// New DB
	quickstartDatabase := client.Database("quickstart")

	// New collection
	podcastsCollection := quickstartDatabase.Collection("podcasts")

	song := modules.Song{
		Title:     "Track01",
		Artist:    "NoName",
		Album:     "Unknown",
		Performer: "Unknown",
	}

	data, err := bson.Marshal(song)

	// Adds an item to the collection
	// you can insert a JSON structure directly without converting it to BSON
	podcastResult, err := podcastsCollection.InsertOne(ctx, data)

	fmt.Println(podcastResult)
	c.JSON(http.StatusOK, podcastResult)

}
