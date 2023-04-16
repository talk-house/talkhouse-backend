package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MG MongoInstance

func ConnectMongo() {

	uri := os.Getenv("DATABASE_URI")

	if uri := os.Getenv("DATABASE_URI"); uri == "" {
		log.Fatal("You must set your 'DATABASE_URI' environmental variable. ")
	}

	// Create a new client and connect to the server
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	db := client.Database("talkhouse")
	MG = MongoInstance{
		Client: client,
		Db:     db,
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged to database.")
}
