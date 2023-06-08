package dto

import (
	"time"
)

// ProjectQueryParam 查询参数
type ProjectQueryParam struct {
	ID int64 `json:"id" form:"id"`
	// 项目 ID
	ProjectName string `json:"project_name" form:"project_name"`
	// 项目 名称
	ProjectPublisherID int64 `json:"project_publisher_id" form:"project_publisher_id"`
	// 项目发布者 ID
	MinProjectBudget int64 `json:"min_project_budget" form:"min_project_budget"`
	// 项目预算最小值，保留两位小数
	MaxProjectBudget int64 `json:"max_project_budget" form:"max_project_budget"`
	// 项目预算最大值，保留两位小数
	ProjectDuration string `json:"project_duration" form:"project_duration"`
	// 项目周期
	CreateTimeBegin time.Time `json:"create_time_begin" form:"create_time_begin"`
	// 创建时间开始
	CreateTimeEnd time.Time `json:"create_time_end" form:"create_time_end"`
	// 创建事件结束
	UpdateTimeBegin time.Time `json:"update_time_begin" form:"update_time_begin"`
	// 最后一次更新时间开始
	UpdateTimeEnd time.Time `json:"update_time_end" form:"update_time_end"`
	// 最后一次更新时间结束
	ProjectScheduleID int64 `json:"project_schedule_id" form:"project_schedule_id"`
	// 项目进度ID
	WorkType int `json:"work_type" form:"work_type"`
	// 工作类型(0为全职兼职都可以 1为全职 2为兼职)
	ProjectPositionTypeList []int64 `json:"project_position_type_list" form:"project_position_type_list"`
	// 项目所需工程师类型列表，只存放 ID
	ProjectTypeList []int64 `json:"project_type_list" form:"project_type_list"`
	// 项目类型列表，只存放 ID
	OrderType int `json:"order_type" form:"order_type"`
	// 排序类型 0为按金额排序 1为按周期排序
	OrderWay int `json:"order_way" form:"order_way"`
	// 排序方式 0为正序 1为逆序
}
