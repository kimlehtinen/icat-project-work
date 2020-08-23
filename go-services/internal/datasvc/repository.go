package datasvc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kim3z/icat-project-work/pkg/dbcontext"
	"github.com/kim3z/icat-project-work/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	AllTemperature() ([]models.Temperature, error)
	AllBloodPressure() ([]models.BloodPressure, error)
	LastInsertedTemperature() (models.Temperature, error)
	FindBloodPressure(id string) (models.BloodPressure, error)
	FindTemperature(id string) (models.Temperature, error)
}

type BloodPressures []models.BloodPressure

// repository
type repository struct {
	db *mongo.Database
}

// InitRepository creates a new blodpressure repository
func InitRepository(db *mongo.Database) Repository {
	return repository{db}
}

// AllBloodPressure returns all bp results
func (r repository) AllBloodPressure() ([]models.BloodPressure, error) {
	collection := r.db.Collection(dbcontext.Collections().BloodPressure)
	results := []models.BloodPressure{}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"createdat", -1}})
	cursor, err := collection.Find(context.TODO(), bson.M{}, findOptions)

	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		var bpResult models.BloodPressure
		err = cursor.Decode(&bpResult)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}

		results = append(results, bpResult)
	}

	return results, nil
}

// AllTemperature returns all bp results
func (r repository) AllTemperature() ([]models.Temperature, error) {
	collection := r.db.Collection(dbcontext.Collections().Temperature)
	results := []models.Temperature{}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"createdat", -1}})
	cursor, err := collection.Find(context.TODO(), bson.M{}, findOptions)

	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		var result models.Temperature
		err = cursor.Decode(&result)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}

		results = append(results, result)
	}

	return results, nil
}

// CurrentTemperature returns a temperature result by id
func (r repository) LastInsertedTemperature() (models.Temperature, error) {
	var temperatureDB models.Temperature
	collection := r.db.Collection(dbcontext.Collections().Temperature)
	//filter := bson.M{} // last inserted
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	findOptions := options.FindOne()
	findOptions.SetSort(bson.D{{"_id", -1}})
	err := collection.FindOne(ctx, bson.D{}, findOptions).Decode(&temperatureDB)

	if err != nil {
		fmt.Println("error retrieving current temperature")
		return models.Temperature{}, err
	}

	return temperatureDB, err
}

// FindBloodPressure returns a bp result by id
func (r repository) FindBloodPressure(id string) (models.BloodPressure, error) {
	var bpDB models.BloodPressure
	objectIDS, _ := primitive.ObjectIDFromHex(id)
	collection := r.db.Collection(dbcontext.Collections().BloodPressure)
	filter := bson.M{"_id": objectIDS}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&bpDB)

	if err != nil {
		fmt.Println("error retrieving bloodpressure with id : " + id)
		return models.BloodPressure{}, err
	}

	return bpDB, err
}

// FindTemperature returns a temperature result by id
func (r repository) FindTemperature(id string) (models.Temperature, error) {
	var temperatureDB models.Temperature
	objectIDS, _ := primitive.ObjectIDFromHex(id)
	collection := r.db.Collection(dbcontext.Collections().Temperature)
	filter := bson.M{"_id": objectIDS}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&temperatureDB)

	if err != nil {
		fmt.Println("error retrieving temperature with id : " + id)
		return models.Temperature{}, err
	}

	return temperatureDB, err
}
