package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoClient is the currently connected mongo client
var MongoClient *mongo.Client

// MongoConnect starts a connection with mongodb
func MongoConnect() {
	MongoClient, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE")))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	MongoClient.Connect(ctx)
	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("could not connect to database")
	}
}
