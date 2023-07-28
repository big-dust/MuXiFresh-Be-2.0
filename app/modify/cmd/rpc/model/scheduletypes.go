package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schedule struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID          primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Name            string             `bson:"name,omitempty" json:"name,omitempty"`
	School          string             `bson:"school,omitempty" json:"school,omitempty"`
	Major           string             `bson:"major,omitempty" json:"major,omitempty"`
	Group           string             `bson:"group,omitempty" json:"group,omitempty"`
	EntryFormStatus string             `bson:"entry_form_status,omitempty" json:"entry_form_status,omitempty"`
	AdmissionStatus string             `bson:"admission_status,omitempty" json:"admission_status,omitempty"`
	UpdateAt        time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt        time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
	
}
