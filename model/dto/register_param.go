package dto

// RegisterParam 用户注册的时候需要提供的信息
// 注册的时候适用的参数校验规则:
// Username: 参数必须传递, 用户名字符串最小长度为 1, 最大长度为 32
// Password: 参数必须传递, 密码字符串最小长度为 1, 最大长度为 32
// Tel: 电话号码必须为数字, 长度必须是 11 位
type RegisterParam struct {
	Username    string `json:"username" form:"username" `        // 用户名
	Password    string `json:"password" form:"password"`         // 密码
	Phone       string `json:"phone" form:"phone" `              // 电话
	Code        string `json:"code" form:"code"`                 // 验证码
	IdentityTag int    `json:"identity_tag" form:"identity_tag"` // 身份标志;1:雇主，2:开发者
}
