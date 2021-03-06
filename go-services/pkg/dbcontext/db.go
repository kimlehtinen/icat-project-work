package dbcontext

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type databaseCollections struct {
	BloodPressure string
	Temperature   string
}

// NewConnection creates new database connection
func NewConnection() (*mongo.Database, error) {
	// Database Config
	clientOptions := options.Client().ApplyURI("mongodb://mongo-db:27017")
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	db := client.Database("icat-project-work")
	return db, nil
}

func Collections() databaseCollections {
	collections := databaseCollections{
		BloodPressure: "bloodpressure",
		Temperature:   "temperature",
	}

	return collections
}
