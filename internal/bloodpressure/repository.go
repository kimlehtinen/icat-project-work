package bloodpressure

import (
	"context"
	"log"

	"github.com/kim3z/icat-project-work/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	All() ([]entity.BloodPressure, error)
}

type BloodPressures []entity.BloodPressure

// repository
type repository struct {
	db *mongo.Database
}

var collectionName = "test"

// InitRepository creates a new blodpressure repository
func InitRepository(db *mongo.Database) Repository {
	return repository{db}
}

func (r repository) All() ([]entity.BloodPressure, error) {
	collection := r.db.Collection(collectionName)
	results := []entity.BloodPressure{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		var bpResult entity.BloodPressure
		err = cursor.Decode(&bpResult)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}

		results = append(results, bpResult)
	}

	return results, nil
}
