package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ ScheduleModel = (*customScheduleModel)(nil)

type (
	// ScheduleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScheduleModel.
	ScheduleModel interface {
		scheduleModel
		FindOneByUserId(ctx context.Context, userId string) (*Schedule, error)
	}

	customScheduleModel struct {
		*defaultScheduleModel
	}
)

// NewScheduleModel returns a model for the mongo.
func NewScheduleModel(url, db, collection string) ScheduleModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customScheduleModel{
		defaultScheduleModel: newDefaultScheduleModel(conn),
	}
}

func (m *customScheduleModel) FindOneByUserId(ctx context.Context, userId string) (*Schedule, error) {
	uid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data Schedule

	err = m.conn.FindOne(ctx, &data, bson.M{"user_id": uid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
