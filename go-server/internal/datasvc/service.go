package datasvc

import (
	"time"

	"github.com/kim3z/icat-project-work/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Find(id string) (models.BloodPressure, error)
	All() ([]models.BloodPressure, error)
	Create(input CreateBloodPressureRequest) (models.BloodPressure, error)
}

type CreateBloodPressureRequest struct {
	Diastolic   float64 `json:"diastolic"`
	Systolic    float64 `json:"systolic"`
	PulsePerMin float64 `json:"pulse_per_min"`
}

type service struct {
	repo Repository
}

// InitService creates a new bp service
func InitService(repo Repository) Service {
	return service{repo}
}

// Find returns a bp result by id
func (s service) Find(id string) (models.BloodPressure, error) {
	bp, err := s.repo.Find(id)
	if err != nil {
		return models.BloodPressure{}, err
	}
	return bp, nil
}

// All returns all bp results
func (s service) All() ([]models.BloodPressure, error) {
	results, err := s.repo.All()
	if err != nil {
		return []models.BloodPressure{}, nil
	}
	return results, nil
}

// Create creates a new bp result
func (s service) Create(req CreateBloodPressureRequest) (models.BloodPressure, error) {
	now := time.Now()
	insertResult, err := s.repo.Create(models.BloodPressure{
		Diastolic:   req.Diastolic,
		Systolic:    req.Systolic,
		PulsePerMin: req.PulsePerMin,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		return models.BloodPressure{}, err
	}
	idStr := insertResult.InsertedID.(primitive.ObjectID).Hex()
	return s.Find(idStr)
}
