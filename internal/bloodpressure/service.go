package bloodpressure

import "github.com/kim3z/icat-project-work/internal/entity"

type Service interface {
	All() ([]entity.BloodPressure, error)
}

type service struct {
	repo Repository
}

func InitService(repo Repository) Service {
	return service{repo}
}

func (s service) All() ([]entity.BloodPressure, error) {
	results, err := s.repo.All()
	if err != nil {
		return []entity.BloodPressure{}, nil
	}
	return results, nil
}
