// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFileType = "file_type"

// FileType mapped from table <file_type>
type FileType struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:文件类型 ID、自增" json:"id"`
	TypeName   string    `gorm:"column:type_name;comment:文件类型名称" json:"type_name"`
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"create_time"`
}

// TableName FileType's table name
func (*FileType) TableName() string {
	return TableNameFileType
}
