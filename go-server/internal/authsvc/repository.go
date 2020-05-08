package authsvc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kim3z/icat-project-work/pkg/jwtauth"

	"github.com/kim3z/icat-project-work/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Find(id string) (models.User, error)
	All() ([]models.User, error)
	Register(user models.User) (*mongo.InsertOneResult, error)
	Auth(user AuthUser) (models.JwtToken, error)
}

type Users []models.User

// repository
type repository struct {
	db *mongo.Database
}

var collectionName = "user"

// InitRepository creates a new blodpressure repository
func InitRepository(db *mongo.Database) Repository {
	return repository{db}
}

// Find finds a user by id
func (r repository) Find(id string) (models.User, error) {
	var userDB models.User
	objectIDS, _ := primitive.ObjectIDFromHex(id)
	collection := r.db.Collection(collectionName)
	filter := bson.M{"_id": objectIDS}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&userDB)

	if err != nil {
		fmt.Println("error retrieving user user by id : " + id)
		return models.User{}, err
	}

	return userDB, err
}

// All returns all users
func (r repository) All() ([]models.User, error) {
	collection := r.db.Collection(collectionName)
	users := []models.User{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		var usersResult models.User
		err = cursor.Decode(&usersResult)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}

		users = append(users, usersResult)
	}

	return users, nil
}

// Create creates a new user
func (r repository) Register(user models.User) (*mongo.InsertOneResult, error) {
	// bcrypt password
	pswHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(pswHash)

	// save user to db
	collection := r.db.Collection(collectionName)
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return &mongo.InsertOneResult{}, err
	}
	return insertResult, nil
}

func (r repository) Auth(user AuthUser) (models.JwtToken, error) {
	var userDB models.User
	collection := r.db.Collection(collectionName)
	filter := bson.M{"email": user.Email}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&userDB)

	if err != nil {
		fmt.Println("error retrieving user user by email : " + user.Email)
		return models.JwtToken{}, err
	}

	// compare given password with db password
	pswErr := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)) // should only return nil if passwords match
	if pswErr != nil || pswErr == bcrypt.ErrMismatchedHashAndPassword {
		return models.JwtToken{}, pswErr
	}

	jwtToken, err := jwtauth.GenerateToken(userDB)

	if err != nil {
		log.Fatal(err)
		return models.JwtToken{}, err
	}

	return jwtToken, nil
}
