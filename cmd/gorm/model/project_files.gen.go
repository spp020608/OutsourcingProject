// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProjectFile = "project_files"

// ProjectFile mapped from table <project_files>
type ProjectFile struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:项目文件 ID、自增" json:"id"`
	ProjectID  int64     `gorm:"column:project_id;comment:项目id" json:"project_id"`
	UserID     int64     `gorm:"column:user_id;comment:用户id" json:"user_id"`
	FileTypeID int64     `gorm:"column:file_type_id;comment:文件类型id" json:"file_type_id"`
	FileName   string    `gorm:"column:file_name;comment:文件原名字" json:"file_name"`
	FileURL    string    `gorm:"column:file_url;comment:文件路径" json:"file_url"`
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName ProjectFile's table name
func (*ProjectFile) TableName() string {
	return TableNameProjectFile
}
