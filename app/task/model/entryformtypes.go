package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EntryForm struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Grade    string             `bson:"grade,omitempty" json:"grade,omitempty"`
	Group    string             `bson:"group,omitempty" json:"group,omitempty"`
	Avatar   string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	College  string             `bson:"college,omitempty" json:"college,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
