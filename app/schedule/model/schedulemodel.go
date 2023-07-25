package model

import "github.com/zeromicro/go-zero/core/stores/mon"

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
func NewScheduleModel(url, db, collection string) ScheduleModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customScheduleModel{
		defaultScheduleModel: newDefaultScheduleModel(conn),
	}
}
