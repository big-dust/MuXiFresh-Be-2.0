package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ UserAuthModel = (*customUserAuthModel)(nil)

type (
	// UserAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthModel.
	UserAuthModel interface {
		userAuthModel
		FindByEmPass(ctx context.Context, Email string, Password string) (*UserAuth, error)
		UpdateByEm(ctx context.Context, data *UserAuth) (*mongo.UpdateResult, error)
	}

	customUserAuthModel struct {
		*defaultUserAuthModel
	}
)

// NewUserAuthModel returns a model for the mongo.
func NewUserAuthModel(url, db, collection string) UserAuthModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customUserAuthModel{
		defaultUserAuthModel: newDefaultUserAuthModel(conn),
	}
}
