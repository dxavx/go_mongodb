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
		v1.GET("/url.insert", InsertUrl)

	}
	router.Run(":8080")

}

func InsertUrl(c *gin.Context) {

	client, ctx, _ := modules.ConnectMongo()
	defer client.Disconnect(ctx)

	// List DB
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	// Новая DB
	quickstartDatabase := client.Database("quickstart")

	// Новая коллекция
	podcastsCollection := quickstartDatabase.Collection("podcasts")
	//episodesCollection := quickstartDatabase.Collection("episodes")

	// Добавляет элемент в коллекцию
	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
	})

	fmt.Println(podcastResult)

	c.JSON(http.StatusOK, podcastResult)

}
