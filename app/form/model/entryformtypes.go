package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EntryForm struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId        primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Avatar        string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	Name          string             `bson:"name,omitempty" json:"name,omitempty"`
	StuNumber     string             `bson:"stuNumber,omitempty" json:"stuNumber,omitempty"`
	School        string             `bson:"school,omitempty" json:"school,omitempty"`
	Major         string             `bson:"major,omitempty" json:"major,omitempty"`
	Grade         string             `bson:"grade,omitempty" json:"grade,omitempty"`
	Gender        string             `bson:"gender,omitempty" json:"gender,omitempty"`
	Email         string             `bson:"email,omitempty" json:"email,omitempty"`
	QQ            string             `bson:"qq,omitempty" json:"qq,omitempty"`
	Phone         string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Group         string             `bson:"group,omitempty" json:"group,omitempty"`
	Reason        string             `bson:"reason,omitempty" json:"reason,omitempty"`
	Knowledge     string             `bson:"knowledge,omitempty" json:"knowledge,omitempty"`
	SelfIntro     string             `bson:"selfIntro,omitempty" json:"selfIntro,omitempty"`
	ExtraQuestion string             `bson:"extraQuestion,omitempty" json:"extraQuestion,omitempty"`
	UpdateAt      time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt      time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
