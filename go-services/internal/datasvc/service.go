package datasvc

import (
	"github.com/kim3z/icat-project-work/pkg/models"
)

type Service interface {
	FindBloodPressure(id string) (models.BloodPressure, error)
	AllBloodPressure() ([]models.BloodPressure, error)
	FindTemperature(id string) (models.Temperature, error)
	AllTemperature() ([]models.Temperature, error)
}

type service struct {
	repo Repository
}

// InitService creates a new bp service
func InitService(repo Repository) Service {
	return service{repo}
}

// FindBloodPressure returns a bp result by id
func (s service) FindBloodPressure(id string) (models.BloodPressure, error) {
	bp, err := s.repo.FindBloodPressure(id)
	if err != nil {
		return models.BloodPressure{}, err
	}
	return bp, nil
}

// AllBloodPressure returns all bp results
func (s service) AllBloodPressure() ([]models.BloodPressure, error) {
	results, err := s.repo.AllBloodPressure()
	if err != nil {
		return []models.BloodPressure{}, nil
	}
	return results, nil
}

// FindTemperature returns a temperature result by id
func (s service) FindTemperature(id string) (models.Temperature, error) {
	temperature, err := s.repo.FindTemperature(id)
	if err != nil {
		return models.Temperature{}, err
	}
	return temperature, nil
}

// AllTemperature returns all temperature results
func (s service) AllTemperature() ([]models.Temperature, error) {
	results, err := s.repo.AllTemperature()
	if err != nil {
		return []models.Temperature{}, nil
	}
	return results, nil
}
