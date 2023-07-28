package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Avatar    string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	NickName  string             `bson:"nickname,omitempty" json:"nickname,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	StudentNo string             `bson:"student_no,omitempty" json:"student_no,omitempty"`
	Group     string             `bson:"group,omitempty" json:"group,omitempty"`
	UserType  string             `bson:"user_type,omitempty" json:"user_type,omitempty"`
	FormID    primitive.ObjectID `bson:"form_id,omitempty" json:"form_id,omitempty"`
	UpdateAt  time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt  time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
