package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schedule struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AdmissionStatus string             `bson:"admission_status,omitempty" json:"admission_status,omitempty"`
	UpdateAt        time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt        time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
