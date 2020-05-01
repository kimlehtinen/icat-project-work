package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type BloodPressure struct {
	ID  primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Foo string             `json:"foo"`
}
