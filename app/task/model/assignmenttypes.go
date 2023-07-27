package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Assignment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TitleText string             `bson:"title,omitempty" json:"title,omitempty"`
	Group     string             `bson:"group,omitempty" json:"group,omitempty"`
	Content   string             `bson:"content,omitempty" json:"content,omitempty"`
	Urls      []string           `bson:"urls,omitempty" json:"urls,omitempty"`
	UpdateAt  time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt  time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
