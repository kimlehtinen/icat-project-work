package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kim3z/icat-project-work/internal/bloodpressure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var serverPort = 8080

type Todo struct {
	Foo string `json:"foo"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// test db
		db := connect()
		collection := db.Collection("test")
		todos := []Todo{}
		// findOptions := options.Find()
		cursor, err := collection.Find(context.TODO(), bson.M{})

		if err != nil {
			panic(err)
		}
		fmt.Println("CURSOR NEXT")
		for cursor.Next(context.TODO()) {
			var todo Todo
			err = cursor.Decode(&todo)
			if err != nil {
				log.Fatal("Error on Decoding the document", err)
			}
			fmt.Printf("%+v\n", todo)
			todos = append(todos, todo)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todos); err != nil {
			panic(err)
		}
	})

	// blood pressure routes
	bloodpressure.RegisterHandlers(router, bloodpressure.InitService(bloodpressure.InitRepository("testdbcontext")))

	fmt.Println("Server listening!")
	http.ListenAndServe(":8080", router)
}

func connect() *mongo.Database {
	// Database Config
	clientOptions := options.Client().ApplyURI("mongodb://mongo_db:27017")
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
	// controllers.TodoCollection(db)
	return db
}
