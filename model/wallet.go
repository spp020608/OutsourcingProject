package model

import "github.com/shopspring/decimal"

type Wallet struct {
	UserID  int             `json:"user_id" gorm:"column:user_id"`     // 用户 ID
	Balance decimal.Decimal `json:"balance" sql:"type:decimal(20,8);"` // 用户钱包余额
}
