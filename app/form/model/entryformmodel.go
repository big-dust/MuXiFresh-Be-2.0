package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var _ EntryFormModel = (*customEntryFormModel)(nil)

type (
	// EntryFormModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEntryFormModel.
	EntryFormModel interface {
		entryFormModel
		InsertReturnID(ctx context.Context, data *EntryForm) (interface{}, error)
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
