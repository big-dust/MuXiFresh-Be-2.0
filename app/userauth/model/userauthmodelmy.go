package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (m *customUserAuthModel) FindByEmPass(ctx context.Context, Email string, Password string) (*UserAuth, error) {
	var data UserAuth

	err := m.conn.FindOne(ctx, &data, bson.M{"email": Email, "password": Password})

	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserAuthModel) UpdateByEm(ctx context.Context, data *UserAuth) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"email": data.Email}, bson.M{"$set": data})

	return res, err
}
