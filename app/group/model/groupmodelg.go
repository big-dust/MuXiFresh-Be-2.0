package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *customGroupModel) FindByName(ctx context.Context, name string) (*Group, error) {
	var data Group
	err := m.conn.FindOne(ctx, &data, bson.M{"name": name})

	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
