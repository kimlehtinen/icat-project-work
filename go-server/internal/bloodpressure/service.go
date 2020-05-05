package bloodpressure

import (
	"time"

	"github.com/kim3z/icat-project-work/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Find(id string) (entity.BloodPressure, error)
	All() ([]entity.BloodPressure, error)
	Create(input CreateBloodPressureRequest) (entity.BloodPressure, error)
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
func (s service) Find(id string) (entity.BloodPressure, error) {
	bp, err := s.repo.Find(id)
	if err != nil {
		return entity.BloodPressure{}, err
	}
	return bp, nil
}

// All returns all bp results
func (s service) All() ([]entity.BloodPressure, error) {
	results, err := s.repo.All()
	if err != nil {
		return []entity.BloodPressure{}, nil
	}
	return results, nil
}

// Create creates a new bp result
func (s service) Create(req CreateBloodPressureRequest) (entity.BloodPressure, error) {
	now := time.Now()
	insertResult, err := s.repo.Create(entity.BloodPressure{
		Diastolic:   req.Diastolic,
		Systolic:    req.Systolic,
		PulsePerMin: req.PulsePerMin,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		return entity.BloodPressure{}, err
	}
	idStr := insertResult.InsertedID.(primitive.ObjectID).Hex()
	return s.Find(idStr)
}
