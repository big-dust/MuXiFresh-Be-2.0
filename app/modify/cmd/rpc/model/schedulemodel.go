package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ ScheduleModel = (*customScheduleModel)(nil)

type (
	// ScheduleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScheduleModel.
	ScheduleModel interface {
		scheduleModel
	}

	customScheduleModel struct {
		*defaultScheduleModel
	}
)

// NewScheduleModel returns a model for the mongo.
func NewScheduleModel(url, db, collection string, c cache.CacheConf) ScheduleModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customScheduleModel{
		defaultScheduleModel: newDefaultScheduleModel(conn),
	}
}
