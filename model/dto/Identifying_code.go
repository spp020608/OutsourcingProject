package dto

type IdentifyingCode struct {
	Phone    string `json:"phone" form:"phone" `      // 电话
	Identity string `json:"identity" form:"identity"` // 身份标志;1:雇主，2:开发者,3:忘记密码
}
