package iotsvc

import (
	"time"

	"github.com/kim3z/icat-project-work/internal/datasvc"

	"github.com/kim3z/icat-project-work/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	StoreBloodPressure(input CreateBloodPressureRequest) (models.BloodPressure, error)
	StoreTemperature(temperature float64) (models.Temperature, error)
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

// StoreBloodPressure creates a new bp result
func (s service) StoreBloodPressure(req CreateBloodPressureRequest) (models.BloodPressure, error) {
	now := time.Now()
	insertResult, err := s.iotRepo.StoreBloodPressure(models.BloodPressure{
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

// StoreBloodPressure creates a new temperature result
func (s service) StoreTemperature(temperature float64) (models.Temperature, error) {
	now := time.Now()
	insertResult, err := s.iotRepo.StoreTemperature(models.Temperature{
		Temperature: temperature,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		return models.Temperature{}, err
	}
	idStr := insertResult.InsertedID.(primitive.ObjectID).Hex()

	return s.dataRepo.FindTemperature(idStr)
}
