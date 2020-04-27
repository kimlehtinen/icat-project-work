package bloodpressure

import "github.com/kim3z/icat-project-work/internal/entity"

type Repository interface {
	All() ([]entity.BloodPressure, error)
}

type BloodPressures []entity.BloodPressure

// repository
type repository struct {
	db string // @TODO: change to a db context
}

// InitRepository creates a new blodpressure repository
func InitRepository(db string) Repository {
	return repository{db}
}

func (r repository) All() ([]entity.BloodPressure, error) {
	results := BloodPressures{
		entity.BloodPressure{Title: "Test from repo"},
		entity.BloodPressure{Title: "Test2"},
	}

	return results, nil
}
