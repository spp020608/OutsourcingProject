// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSchedule = "schedule"

// Schedule mapped from table <schedule>
type Schedule struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:进度 ID、自增" json:"id"`
	ScheduleName string    `gorm:"column:schedule_name;not null;comment:进度名称" json:"schedule_name"`
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName Schedule's table name
func (*Schedule) TableName() string {
	return TableNameSchedule
}
