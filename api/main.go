package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"time"
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

	// Подключение с авторизацией
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017").SetAuth(options.Credential{
		AuthMechanism:           "",
		AuthMechanismProperties: nil,
		AuthSource:              "admin",
		Username:                "root",
		Password:                "example",
		PasswordSet:             false,
	}))

	if err != nil {
		log.Fatal("Error connect MongoDB")
	}

	// Таймаут подключения
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal()
	}

	defer client.Disconnect(ctx)

	// Ping DB
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

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
