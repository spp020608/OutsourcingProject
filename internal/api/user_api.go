package api

import (
	"Outsourcing/cmd/gorm/model"
	"Outsourcing/internal/pkg/resp"
	"Outsourcing/internal/service"
	"Outsourcing/model/dto"
	"Outsourcing/pkg/validate"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var UserApi = new(userApi)

type userApi struct{}

// SendCode Post 请求, 发送 Json 数据, 参数在消息体中
//
//	@Summary		Post 请求, 发送 Json 数据, 参数在消息体中
//	@Description	Post 请求, 发送 Json 数据, 参数在消息体中
//	@Tags			userApi
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.IdentifyingCode	true	"前端发送的 Json 对象"
//	@Success		200		{object}	swagger.HttpOk		"OK"
//	@Failure		400		{object}	swagger.Http400		"Bad Request"
//	@Failure		404		{object}	swagger.Http404		"Page Not Found"
//	@Failure		500		{object}	swagger.Http500		"InternalError"
//	@Router			/user/sendCode [post]
func (userApi) SendCode(c *gin.Context) {
	var u dto.IdentifyingCode
	if err := c.ShouldBind(&u); err != nil {
		trans := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", trans)
		resp.InternalServerError(c, trans)
		return
	}
	code := service.UserService.SendCode(&u)
	if code == 0 {
		zap.S().Errorf("发送验证码失败\n")
		resp.InternalServerError(c, "发送验证码失败")
		return
	}
	resp.Ok(c, "发送验证码成功", nil)
}

// Register 用户注册
//
//	@Summary		Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Description	Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Tags			userApi
//	@Accept			json
//	@Produce		json
//	@Param			registerParam	body		dto.RegisterParam	true	"前端发送的 Json 对象"//	@Success	200	{object}	swagger.HttpOk	"OK"
//	@Failure		400				{object}	swagger.Http400		"Bad Request"
//	@Failure		404				{object}	swagger.Http404		"Page Not Found"
//	@Failure		500				{object}	swagger.Http500		"InternalError"
//	@Router			/user/register [post]
func (userApi) Register(c *gin.Context) {

	var registerParam dto.RegisterParam
	// 将请求数据绑定到结构体
	if err := c.ShouldBind(&registerParam); err != nil {
		trans := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", trans)
		resp.InternalServerError(c, trans)
		return
	}

	register, err := service.UserService.Register(&registerParam)
	if err != nil {
		zap.S().Errorf("注册用户失败: %s\n", err)
		resp.InternalServerError(c, "注册用户失败")
		return
	}
	resp.Ok(c, register, nil)

}

// Login 用户登录
//
//	@Summary		Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Description	Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Tags			userApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			username	formData	string			true	"要注册的用户名"	minlength(1)	maxlength(20)
//	@Param			password	formData	string			true	"用户密码"		minlength(1)	maxlength(128)
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/user/login [post]
func (userApi) Login(c *gin.Context) {
	var loginInfo dto.LoginInfo
	if err := c.ShouldBind(&loginInfo); err != nil {
		// 翻译错误消息
		translatedErr := validate.TranslateError(err)
		// 向前端回写错误数据
		resp.InternalServerError(c, translatedErr)
		// 将错误记录到日志文件
		zap.S().Errorf("用户登录失败: %s\n", translatedErr)
		return
	}

	fmt.Printf("%#v\n", loginInfo)

	// 调用 service 层的登录函数
	token, err := service.UserService.Login(&loginInfo)

	//根据前端传回来的string判断是否未注册
	if token == "用户名不存在,请先注册" {
		resp.Unauthorized(c, "用户请先注册")
		return
	}

	if err != nil {
		// 向前端回写错误数据
		resp.InternalServerError(c, err.Error())
		// 将错误记录到日志文件
		zap.S().Errorf("用户登录失败: %s\n", err.Error())
		return
	}

	if err != nil {
		resp.InternalServerError(c, "请求失败")
		zap.S().Errorf("请求失败：%s\n", err.Error())
		return
	}

	resp.Ok(c, "登录成功", token)

}

// TestAuth 访问该接口的时候, 用户必须是已经登录成功
//
//	@Summary		访问该接口的时候, 用户必须是已经登录成功
//	@Description	访问该接口的时候, 用户必须是已经登录成功
//	@Tags			userApi
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	swagger.HttpOk	"OK"
//	@Failure		400	{object}	swagger.Http400	"Bad Request"
//	@Failure		404	{object}	swagger.Http404	"Page Not Found"
//	@Failure		500	{object}	swagger.Http500	"InternalError"
//	@Router			/user/testAuth [get]
//	@Security		ApiKeyAuth
func (userApi) TestAuth(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		zap.L().Info("从当前上下文中获取登录用户失败")
		return
	}
	zap.S().Info(user)

	c.ShouldBindJSON(&user)
	userModel, ok := user.(*model.User)
	if !ok {
		zap.L().Info("将 interface{} 转换成 User 失败")
		return
	}
	zap.L().Info("从上下文中获取用户信息成功",
		zap.Int("id", int(userModel.ID)),
		zap.String("username", userModel.Username),
		zap.String("password", userModel.Password),
	)
}
