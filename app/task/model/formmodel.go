package model

import (
	"MuXiFresh-Be-2.0/app/userauth/model"
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ FormModel = (*customFormModel)(nil)

type (
	// FormModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFormModel.
	FormModel interface {
		formModel
		FindOneByUserId(ctx context.Context, userId string) (*Form, error)
	}

	customFormModel struct {
		*defaultFormModel
	}
)

// NewFormModel returns a model for the mongo.
func NewFormModel(url, db, collection string) FormModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customFormModel{
		defaultFormModel: newDefaultFormModel(conn),
	}
}

func (m *customFormModel) FindOneByUserId(ctx context.Context, userId string) (*Form, error) {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, model.ErrInvalidObjectId
	}

	var data Form

	err = m.conn.FindOne(ctx, &data, bson.M{"user_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}
