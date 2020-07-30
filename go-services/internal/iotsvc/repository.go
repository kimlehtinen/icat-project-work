package iotsvc

import (
	"context"
	"log"

	"github.com/kim3z/icat-project-work/pkg/dbcontext"

	"github.com/kim3z/icat-project-work/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	StoreBloodPressure(models.BloodPressure) (*mongo.InsertOneResult, error)
	StoreTemperature(models.Temperature) (*mongo.InsertOneResult, error)
}

// repository
type repository struct {
	db *mongo.Database
}

// InitRepository creates a new blodpressure repository
func InitRepository(db *mongo.Database) Repository {
	return repository{db}
}

// StoreBloodPressure creates a new bp result
func (r repository) StoreBloodPressure(bp models.BloodPressure) (*mongo.InsertOneResult, error) {
	collection := r.db.Collection(dbcontext.Collections().BloodPressure)
	insertResult, err := collection.InsertOne(context.TODO(), bp)
	if err != nil {
		log.Fatal(err)
		return &mongo.InsertOneResult{}, err
	}
	return insertResult, nil
}

// StoreBloodPressure creates a new temperature result
func (r repository) StoreTemperature(temperature models.Temperature) (*mongo.InsertOneResult, error) {
	collection := r.db.Collection(dbcontext.Collections().Temperature)
	insertResult, err := collection.InsertOne(context.TODO(), temperature)
	if err != nil {
		log.Fatal(err)
		return &mongo.InsertOneResult{}, err
	}
	return insertResult, nil
}
