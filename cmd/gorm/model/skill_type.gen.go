// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSkillType = "skill_type"

// SkillType mapped from table <skill_type>
type SkillType struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SkillName    string    `gorm:"column:skill_name;comment:技能名字" json:"skill_name"`
	SkillExplain string    `gorm:"column:skill_explain;comment:技能描述" json:"skill_explain"`
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName SkillType's table name
func (*SkillType) TableName() string {
	return TableNameSkillType
}