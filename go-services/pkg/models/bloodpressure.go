package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BloodPressure struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Diastolic   float64            `json:"diastolic"`
	Systolic    float64            `json:"systolic"`
	PulsePerMin float64            `json:"pulse_per_min"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
