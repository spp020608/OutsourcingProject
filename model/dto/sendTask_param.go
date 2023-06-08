package dto

import (
	"github.com/shopspring/decimal"
	"mime/multipart"
	"time"
)

type ProjectPublishParam struct {
	//项目ID
	ProjectId int64 `json:"project_id" form:"project_id"`
	//项目名称
	ProjectName string `json:"project_name" form:"project_name"`
	//项目预算，保留两位小数
	ProjectBudget decimal.Decimal `json:"project_budget" form:"project_budget"`
	//项目简介
	ProjectShortIntro string `json:"project_short_intro" form:"project_short_intro"`
	//项目logo
	ProjectLogo string `json:"project_logo" form:"project_logo"`
	//项目介绍
	ProjectIntroduction string `json:"project_introduction" form:"project_introduction"`
	//项目周期
	ProjectDuration string `json:"project_duration" form:"project_duration"`
	//项目开始时间
	StartTime time.Time `json:"start_time" form:"start_time"`
	//项目结束时间
	EndTime time.Time `json:"end_time" form:"end_time"`
	//工作内容说明
	WorkExplain string `json:"work_explain" form:"work_explain"`
	//工作类型说明
	ProjectTypeIdList []int64 `json:"project_type_id_list" form:"project_type_id_list"`
	//工程师类型ID列表
	EngineerTypeIdList []int64 `json:"engineer_type_id_list" form:"engineer_type_id_list"`
	//企业发布还是个人发布
	IsCompanyPublish int `json:"is_company_publish" form:"is_company_publish"`
	//工作类型
	WorkType int `json:"work_type" form:"work_type"`
	//项目文件
	FileList []*multipart.FileHeader `json:"file_list" form:"file_list"`
}
