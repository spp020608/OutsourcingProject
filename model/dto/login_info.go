package dto

type LoginInfo struct {
	Username string `json:"username" form:"username" biding:"required,min=1,max=32"`
	Password string `json:"password" form:"password" biding:"required,min=1,max=32"`
}
