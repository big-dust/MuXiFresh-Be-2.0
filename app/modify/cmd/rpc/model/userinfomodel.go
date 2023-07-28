package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

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
func NewUserInfoModel(url, db, collection string, c cache.CacheConf) UserInfoModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customUserInfoModel{
		defaultUserInfoModel: newDefaultUserInfoModel(conn),
	}
}
