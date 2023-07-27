package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ UserInfoModel = (*customUserInfoModel)(nil)

type (
	// UserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserInfoModel.
	UserInfoModel interface {
		userInfoModel
		FindByStudentID(ctx context.Context, studentID string) (*UserInfo, error)
	}

	customUserInfoModel struct {
		*defaultUserInfoModel
	}
)

// NewUserInfoModel returns a model for the mongo.
func NewUserInfoModel(url, db, collection string) UserInfoModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customUserInfoModel{
		defaultUserInfoModel: newDefaultUserInfoModel(conn),
	}
}

func (m *defaultUserInfoModel) FindByStudentID(ctx context.Context, studentID string) (*UserInfo, error) {
	var data UserInfo

	err := m.conn.FindOne(ctx, &data, bson.M{"student_no": studentID})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
