package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ UserInfoModel = (*customUserInfoModel)(nil)

type (
	// UserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserInfoModel.
	UserInfoModel interface {
		userInfoModel
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
