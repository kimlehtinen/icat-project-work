package bloodpressure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kim3z/icat-project-work/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Find(id string) (entity.BloodPressure, error)
	All() ([]entity.BloodPressure, error)
	Create(entity.BloodPressure) (*mongo.InsertOneResult, error)
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

func (r repository) Find(id string) (entity.BloodPressure, error) {
	var bpDB entity.BloodPressure
	objectIDS, _ := primitive.ObjectIDFromHex(id)
	collection := r.db.Collection(collectionName)
	filter := bson.M{"_id": objectIDS}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&bpDB)

	if err != nil {
		fmt.Println("error retrieving user userid : " + id)
		return entity.BloodPressure{}, err
	}

	return bpDB, err
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

func (r repository) Create(bp entity.BloodPressure) (*mongo.InsertOneResult, error) {
	collection := r.db.Collection(collectionName)
	insertResult, err := collection.InsertOne(context.TODO(), bp)
	if err != nil {
		log.Fatal(err)
		return &mongo.InsertOneResult{}, err
	}
	return insertResult, nil
}
