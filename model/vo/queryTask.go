package vo

import (
	"Outsourcing/cmd/gorm/model"
	"Outsourcing/pkg/ctime"
)

type QueryTaskVo struct {
	List      []Item `json:"list"`
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
	Total     int    `json:"total"`
	TotalPage int    `json:"totalPage"`
}

type Item struct {
	ID                   int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:用户 ID、自增" json:"id" form:"id" form:"id"`
	ProjectPublisherID   int64   `gorm:"column:project_publisher_id;not null;comment:项目发布者id" json:"project_publisher_id" form:"project_publisher_id"`
	ProjectPublisherName string  `gorm:"column:project_publisher_name;not null;comment:用户名" json:"project_publisher_name" form:"project_publisher_name"`
	Phone                string  `gorm:"column:phone;comment:手机号" json:"phone" form:"phone"`
	ProjectName          string  `gorm:"column:project_name;not null;comment:项目名字" json:"project_name" form:"project_name"`
	ProjectLogo          string  `gorm:"column:project_logo;comment:项目logo" json:"project_logo" form:"project_logo"`
	ProjectBudget        float64 `gorm:"column:project_budget;default:0.00;comment:项目预算，保留两位小数，单位是元" json:"project_budget" form:"project_budget"`
	ProjectShortIntro    string  `gorm:"column:project_short_intro;comment:项目简介" json:"project_short_intro" form:"project_short_intro"`
	ProjectIntroduction  string  `gorm:"column:project_introduction;not null;comment:项目介绍" json:"project_introduction" form:"project_introduction"`
	ProjectDuration      string  `gorm:"column:project_duration;comment:项目周期，单位是天" json:"project_duration" form:"project_duration"`
	WorkExplain          string  `gorm:"column:work_explain;comment:工作内容说明" json:"work_explain" form:"work_explain"`
	WorkType             int32   `gorm:"column:work_type;comment:工作类型（0为全职，1为兼职，2为全职兼职都可）" json:"work_type" form:"work_type"`
	//CollectCount         int     `json:"collect_count" form:"collect_count"`
	//IsCollect            int32   `gorm:"column:is_collect;comment:是否收藏（1收藏，0未收藏），默认0" json:"is_collect" form:"is_collect"`
	//IsDeveloper          int32   `gorm:"column:is_developer;not null;comment:接单成功(1成功，0未成功，2为拒绝接单)，默认0" json:"is_developer" form:"is_developer"`
	//EndDateWarn          int     `json:"end_date_warn" form:"end_date_warn"`
	//过期提醒 time.now 与end time 比较吧
	CreateTime          ctime.CTime `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"create_time" form:"create_time"`
	UpdateTime          ctime.CTime `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:记录更新时间" json:"update_time" form:"update_time"`
	ProjectScheduleID   int         `gorm:"column:project_schedule_id;not null;comment:状态id"json:"project_schedule_id" form:"project_schedule_id"`
	ProjectScheduleName string      `gorm:"column:project_schedule_name;not null;comment:进度名称"json:"project_schedule_name" form:"project_schedule_name"`

	ProjectPositionTypeList []model.PositionType `json:"project_position_type_list" form:"project_position_type_list"`
	ProjectTypeList         []model.ProjectType  `json:"project_type_list" form:"project_type_list"`
	EngineerList            []Engineer           `json:"engineer_list" form:"engineer_list"`
	ProjectScheduleList     []model.Schedule     `json:"project_schedule_list" form:"project_schedule_list"`
	ProjectFileList         []model.ProjectFile  `json:"project_file_list" form:"project_file_list"`
}

type Engineer struct {
	Address              string                     `json:"address"`
	Balance              int                        `json:"balance"`
	CollegeName          string                     `json:"collegeName"`
	CreateTime           string                     `json:"createTime"`
	Icon                 string                     `json:"icon"`
	IDCard               string                     `json:"idCard"`
	IDCardBack           string                     `json:"idCardBack"`
	IDCardFront          string                     `json:"idCardFront"`
	IDCardFrontImage     string                     `json:"idCardFrontImage"`
	IDCardNo             string                     `json:"idCardNo"`
	IDCardReverseImage   string                     `json:"idCardReverseImage"`
	IfTeam               string                     `json:"ifTeam"`
	PersonalProfile      string                     `json:"personalProfile"`
	Phone                string                     `json:"phone"`
	PlatformProjectList  []model.Project            `json:"platformProjectList"`
	PositionTypeList     []model.PositionType       `json:"positionTypeList"`
	ProjectCount         int                        `json:"projectCount"`
	RealNameApproved     int                        `json:"realNameApproved"`
	Salary               int                        `json:"salary"`
	ShortPersonalProfile string                     `json:"shortPersonalProfile"`
	SkillTypeList        []model.SkillType          `json:"skillTypeList"`
	StarCount            int                        `json:"starCount"`
	UserEducations       []model.UserEducation      `json:"userEducations"`
	UserID               int                        `json:"userId"`
	UserName             string                     `json:"userName"`
	UserProjectList      []model.ProjectNotPlatform `json:"userProjectList"`
	UserWorksList        []model.UserWork           `json:"userWorksList"`
	Wechat               string                     `json:"wechat"`
	WorkStatus           string                     `json:"workStatus"`
	WorkType             int                        `json:"workType"`
	WorkYears            int                        `json:"workYears"`
}
