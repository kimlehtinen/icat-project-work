package iotsvc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kim3z/icat-project-work/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Find(id string) (models.BloodPressure, error)
	Create(models.BloodPressure) (*mongo.InsertOneResult, error)
}

type BloodPressures []models.BloodPressure

// repository
type repository struct {
	db *mongo.Database
}

var collectionName = "test"

// InitRepository creates a new blodpressure repository
func InitRepository(db *mongo.Database) Repository {
	return repository{db}
}

// Find returns a bp result by id
func (r repository) Find(id string) (models.BloodPressure, error) {
	var bpDB models.BloodPressure
	objectIDS, _ := primitive.ObjectIDFromHex(id)
	collection := r.db.Collection(collectionName)
	filter := bson.M{"_id": objectIDS}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&bpDB)

	if err != nil {
		fmt.Println("error retrieving user userid : " + id)
		return models.BloodPressure{}, err
	}

	return bpDB, err
}

// Create creates a new bp result
func (r repository) Create(bp models.BloodPressure) (*mongo.InsertOneResult, error) {
	collection := r.db.Collection(collectionName)
	insertResult, err := collection.InsertOne(context.TODO(), bp)
	if err != nil {
		log.Fatal(err)
		return &mongo.InsertOneResult{}, err
	}
	return insertResult, nil
}
