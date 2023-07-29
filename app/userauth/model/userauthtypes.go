package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAuth struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email      string             `bson:"email,omitempty" json:"email,omitempty"`
	Password   string             `bson:"password,omitempty" json:"password,omitempty"`
	UserInfoID primitive.ObjectID `json:"userInfoID,omitempty" bson:"userInfoID,omitempty"`
	UpdateAt   time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt   time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
