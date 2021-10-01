package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(uri, dbName string) (*mongo.Database, error) {
	// Create a Client to a MongoDB server and use Ping to verify that the
	// server is running.
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to MongoDb %s\n", os.Getenv("MONGO_DBNAME"))
	db := client.Database(dbName)

	return db, nil
}

func CloseMongoDB(mongoDb *mongo.Database) {
	if err := mongoDb.Client().Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
