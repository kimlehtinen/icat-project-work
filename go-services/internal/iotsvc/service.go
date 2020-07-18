package iotsvc

import (
	"time"

	"github.com/kim3z/icat-project-work/internal/datasvc"

	"github.com/kim3z/icat-project-work/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Create(input CreateBloodPressureRequest) (models.BloodPressure, error)
}

type CreateBloodPressureRequest struct {
	Diastolic   float64 `json:"diastolic"`
	Systolic    float64 `json:"systolic"`
	PulsePerMin float64 `json:"pulse_per_min"`
}

type service struct {
	iotRepo  Repository
	dataRepo datasvc.Repository
}

// InitService creates a new bp service
func InitService(iotRepo Repository, dataRepo datasvc.Repository) Service {
	return service{iotRepo, dataRepo}
}

// Create creates a new bp result
func (s service) Create(req CreateBloodPressureRequest) (models.BloodPressure, error) {
	now := time.Now()
	insertResult, err := s.iotRepo.Create(models.BloodPressure{
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

	return s.dataRepo.FindBloodPressure(idStr)
}
