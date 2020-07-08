package modules

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func InsertDB(song Song) interface{} {

	client, ctx, _ := ConnectMongo()
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

	data, err := bson.Marshal(song)

	// Adds an item to the collection
	// you can insert a JSON structure directly without converting it to BSON
	podcastResult, err := podcastsCollection.InsertOne(ctx, data)

	fmt.Println(podcastResult.InsertedID)
	return podcastResult
}
