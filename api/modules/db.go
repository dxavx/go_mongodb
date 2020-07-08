package modules

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func ConnectMongo() (*mongo.Client, context.Context, error) {

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

	// Ping DB
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, err

}
