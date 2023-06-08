package dao

import (
	"Outsourcing/model"
	"gorm.io/gorm"
)

var WalletDao = new(walletDao)

type walletDao struct{}

const walletTableName = "wallet"

// Insert 添加钱包信息
//
//	@parameter	wallet 钱包对象
//	@return		rowsAffected 添加成功返回 1 否则返回 0
//	@return		err 有错误发生
func (walletDao) Insert(db *gorm.DB, wallet *model.Wallet) (rowsAffected int64, err error) {
	tx := db.Table(walletTableName).Create(wallet)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (walletDao) GetAllWalletInfo(db *gorm.DB) (walletList []model.Wallet, err error) {
	walletList = make([]model.Wallet, 100)
	tx := db.Table(walletTableName).Find(&walletList)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return walletList, nil
}
