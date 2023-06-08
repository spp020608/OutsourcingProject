package api

import (
	"Outsourcing/cmd/gorm/model"
	"Outsourcing/internal/pkg/resp"
	"Outsourcing/internal/service"
	"Outsourcing/model/dto"
	"Outsourcing/pkg/validate"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var ProjectApi = new(projectApi)

type projectApi struct{}

// SendTask Post 请求, 发送 Json 数据, 参数在消息体中
//
//	@Summary		Post 请求, 发送 Json 数据, 参数在消息体中
//	@Description	Post 请求, 发送 Json 数据, 参数在消息体中
//	@Tags			projectApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			project_name		formData	string			true	"项目名称"
//	@Param			project_budget		formData	int32			true	"项目预算"
//	@Param			project_short_intro	formData	string			true	"项目简介"
//	@Param			ProjectLogo			formData	string			true	"项目logo"
//	@Param			ProjectIntroduction	formData	string			true	"项目介绍"
//	@Param			ProjectDuration		formData	string			true	"项目周期"
//	@Param			StartTime			formData	string			true	"项目开始时间"
//	@Param			EndTime				formData	string			true	"项目结束时间"
//	@Param			WorkExplain			formData	string			true	"工作内容说明"
//	@Param			ProjectTypeIdList	formData	array			true	"工作类型说明"
//	@Param			EngineerTypeIdList	formData	array			true	"工程师类型ID列表"
//	@Param			WorkType			formData	int32			true	"工作类型"
//	@Param			FileList			formData	file			true	"项目文件"
//	@Success		200					{object}	swagger.HttpOk	"OK"
//	@Failure		400					{object}	swagger.Http400	"Bad Request"
//	@Failure		404					{object}	swagger.Http404	"Page Not Found"
//	@Failure		500					{object}	swagger.Http500	"InternalError"
//	@Router			/project/sendTask [post]
func (projectApi) SendTask(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		zap.L().Info("从当前上下文中获取登录用户失败")
		return
	}
	zap.S().Info(user)

	userModel, ok := user.(*model.User)
	if !ok {
		zap.L().Info("将 interface{} 转换成 User 失败")
		return
	}
	var u dto.ProjectPublishParam
	if err := c.ShouldBind(&u); err != nil {
		trans := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", trans)
		resp.InternalServerError(c, trans)
		return
	}
	result, err := service.ProjectService.SendTask(&u, userModel)
	if result == 0 {
		zap.S().Error("发布任务失败\n", err)
		resp.InternalServerError(c, "发布任务失败")
		return
	}
	resp.Ok(c, "发布任务成功", nil)
}

// QueryTask Post 请求, 发送 Json 数据, 参数在消息体中
//
//	@Summary		Post 请求, 发送表单数据, 参数在消息体中
//	@Description	Post 请求, 发送表单数据, 参数在消息体中
//	@Tags			projectApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			id							formData	int64			false	"用户 ID"
//	@Param			project_name				formData	string			false	"项目 名称"
//	@Param			project_publisher_id		formData	int64			false	"项目发布者id"
//	@Param			min_project_budget			formData	int64			false	"项目预算最小值"
//	@Param			max_project_budget			formData	int64			false	"项目预算最大值"
//	@Param			project_duration			formData	int				false	"项目周期"
//	@Param			create_time_begin			formData	string			false	"创建时间开始"
//	@Param			create_time_end				formData	string			false	"创建事件结束"
//	@Param			update_time_begin			formData	string			false	"最后一次更新时间开始"
//	@Param			update_time_end				formData	string			false	"最后一次更新时间结束"
//	@Param			project_schedule_id			formData	int64			false	"项目进度ID"
//	@Param			work_type					formData	int				false	"工作类型"
//	@Param			project_position_type_list	formData	array			false	"项目所需工程师类型列表"
//	@Param			project_type_list			formData	array			false	"项目类型列表"
//	@Param			order_type					formData	int				false	"排序类型 0为按金额排序 1为按周期排序"
//	@Param			order_way					formData	int				false	"排序方式 0为正序 1为逆序"
//	@Success		200							{object}	swagger.HttpOk	"OK"
//	@Failure		400							{object}	swagger.Http400	"Bad Request"
//	@Failure		404							{object}	swagger.Http404	"Page Not Found"
//	@Failure		500							{object}	swagger.Http500	"InternalError"
//	@Router			/project/queryTask [post]
func (projectApi) QueryTask(c *gin.Context) {

	var u dto.ProjectQueryParam
	if err := c.ShouldBind(&u); err != nil {
		trans := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", trans)
		resp.InternalServerError(c, trans)
		return
	}
	queryTask := service.ProjectService.QueryTask(&u)
	if queryTask == nil {
		zap.S().Error("查询失败\n", nil)
		resp.InternalServerError(c, "查询失败")
		return
	}
	resp.Ok(c, "查询成功", queryTask)

}
