package authsvc

import (
	"time"

	"github.com/kim3z/icat-project-work/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Find(id string) (models.User, error)
	All() ([]models.User, error)
	Register(user AuthUser) (models.User, error)
	Auth(user AuthUser) (models.JwtToken, error)
}

type AuthUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type service struct {
	repo Repository
}

// InitService creates a new user service
func InitService(repo Repository) Service {
	return service{repo}
}

// Auth authenticates user
func (s service) Auth(req AuthUser) (models.JwtToken, error) {
	jwtToken, err := s.repo.Auth(req)
	if err != nil {
		return models.JwtToken{}, err
	}
	return jwtToken, nil
}

// Find returns a user result by id
func (s service) Find(id string) (models.User, error) {
	user, err := s.repo.Find(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// All returns all users
func (s service) All() ([]models.User, error) {
	users, err := s.repo.All()
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

// Create creates a new user
func (s service) Register(req AuthUser) (models.User, error) {
	now := time.Now()
	insertResult, err := s.repo.Register(models.User{
		Email:     req.Email,
		Password:  req.Password,
		Role:      models.NORMAL_USER,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return models.User{}, err
	}
	idStr := insertResult.InsertedID.(primitive.ObjectID).Hex()
	return s.Find(idStr)
}
