// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProjectTypeRelation = "project_type_relation"

// ProjectTypeRelation mapped from table <project_type_relation>
type ProjectTypeRelation struct {
	ID            int64 `gorm:"column:id;primaryKey;autoIncrement:true;comment:关联表自增ID" json:"id"`
	ProjectID     int64 `gorm:"column:project_id;comment:项目ID，对应project表中的id" json:"project_id"`
	ProjectTypeID int64 `gorm:"column:project_type_id;comment:项目类型ID，对应project_type表中的id" json:"project_type_id"`
}

// TableName ProjectTypeRelation's table name
func (*ProjectTypeRelation) TableName() string {
	return TableNameProjectTypeRelation
}