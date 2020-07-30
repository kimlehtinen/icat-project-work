package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Temperature struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Temperature float64            `json:"temperature"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
