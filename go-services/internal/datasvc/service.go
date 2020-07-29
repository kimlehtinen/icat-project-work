package datasvc

import (
	"github.com/kim3z/icat-project-work/pkg/models"
)

type Service interface {
	FindBloodPressure(id string) (models.BloodPressure, error)
	All() ([]models.BloodPressure, error)
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

// All returns all bp results
func (s service) All() ([]models.BloodPressure, error) {
	results, err := s.repo.All()
	if err != nil {
		return []models.BloodPressure{}, nil
	}
	return results, nil
}
