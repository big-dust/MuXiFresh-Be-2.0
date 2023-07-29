package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var _ EntryFormModel = (*customEntryFormModel)(nil)

type (
	// EntryFormModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEntryFormModel.
	EntryFormModel interface {
		entryFormModel
		InsertReturnID(ctx context.Context, data *EntryForm) (interface{}, error)
		FindOneByUserId(ctx context.Context, userId string) (*EntryForm, error)
		FindByGroup(ctx context.Context, group string, startDate time.Time, endDate time.Time, limit int64, offset int64) ([]*EntryForm, error)
	}

	customEntryFormModel struct {
		*defaultEntryFormModel
	}
)

// NewEntryFormModel returns a model for the mongo.
func NewEntryFormModel(url, db, collection string) EntryFormModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customEntryFormModel{
		defaultEntryFormModel: newDefaultEntryFormModel(conn),
	}
}

func (m *defaultEntryFormModel) InsertReturnID(ctx context.Context, data *EntryForm) (interface{}, error) {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	id, err := m.conn.InsertOne(ctx, data)
	return id.InsertedID, err
}

func (m *customEntryFormModel) FindOneByUserId(ctx context.Context, userId string) (*EntryForm, error) {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data EntryForm

	err = m.conn.FindOne(ctx, &data, bson.M{"user_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customEntryFormModel) FindByGroup(ctx context.Context, group string, startDate time.Time, endDate time.Time, limit int64, offset int64) ([]*EntryForm, error) {
	var entryForms []*EntryForm
	err := m.conn.Find(ctx, &entryForms, bson.M{
		"group": group,
		"createdAt": bson.M{
			"$gte": startDate, // 大于等于起始时间
			"$lte": endDate,   // 小于等于结束时间
		},
	}, options.Find().SetSkip(offset).SetLimit(limit).SetSort(bson.D{{"name", 1}}))
	if err != nil {
		return nil, err
	}
	return entryForms, nil
}
