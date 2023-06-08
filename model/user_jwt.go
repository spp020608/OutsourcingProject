package model

import "time"

type UserJwt struct {
	ID                   int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:用户 ID、自增" json:"id"`
	Username             string    `gorm:"column:username;not null;comment:用户名" json:"username"`
	RealName             string    `gorm:"column:real_name;comment:真实姓名" json:"real_name"`
	Sex                  int32     `gorm:"column:sex;comment:性别(1男，0女)" json:"sex"`
	College              string    `gorm:"column:college;comment:毕业院校" json:"college"`
	Phone                string    `gorm:"column:phone;comment:手机号" json:"phone"`
	WorkYears            int32     `gorm:"column:work_years;comment:工作年限（单位是年）" json:"work_years"`
	Salary               int32     `gorm:"column:salary;comment:工资，单位（元/天）" json:"salary"`
	Address              string    `gorm:"column:address;comment:地址" json:"address"`
	Email                string    `gorm:"column:email;comment:电子邮箱" json:"email"`
	Password             string    `gorm:"column:password;comment:密码" json:"password"`
	Wechat               string    `gorm:"column:wechat;comment:微信" json:"wechat"`
	Alipay               string    `gorm:"column:alipay;comment:支付宝" json:"alipay"`
	Qq                   string    `gorm:"column:qq;comment:QQ" json:"qq"`
	IDCard               string    `gorm:"column:id_card;comment:身份证号" json:"id_card"`
	WorkType             int32     `gorm:"column:work_type;comment:工作类型（0为全职，1为兼职，2为全职兼职都可）" json:"work_type"`
	WorkStateID          int64     `gorm:"column:work_state_id;comment:工作状态id" json:"work_state_id"`
	PersonalProfile      string    `gorm:"column:personal_profile;comment:自我介绍" json:"personal_profile"`
	ShortPersonalProfile string    `gorm:"column:short_personal_profile;comment:个人简介" json:"short_personal_profile"`
	CreateTime           time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"create_time"`
	UpdateTime           time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:记录更新时间" json:"update_time"`
	IfStudent            int32     `gorm:"column:if_student;default:2;comment:是不是学生，是为1，不是为2" json:"if_student"`
	AuthType             int32     `gorm:"column:auth_type;comment:认证类型(1为个人认证，2为企业认证,0为未认证)" json:"auth_type"`
	RealNameApproved     int32     `gorm:"column:real_name_approved;comment:实名认证是否通过，通过为1，不通过为2，通过才可接项目" json:"real_name_approved"`
	IsPublisher          int32     `gorm:"column:is_publisher;comment:是发布者为1，是开发者为2" json:"is_publisher"`
	Approver             int64     `gorm:"column:approver;comment:审核人id" json:"approver"`
}
