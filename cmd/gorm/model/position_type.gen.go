// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePositionType = "position_type"

// PositionType mapped from table <position_type>
type PositionType struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:职位类型 ID、自增 ID" json:"id"`
	PositionName    string    `gorm:"column:position_name;not null;comment:职位名称" json:"position_name"`
	PositionExplain string    `gorm:"column:position_explain;comment:职位描述" json:"position_explain"`
	CreateTime      time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"create_time"`
}

// TableName PositionType's table name
func (*PositionType) TableName() string {
	return TableNamePositionType
}
