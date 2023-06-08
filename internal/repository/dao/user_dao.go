package dao

import (
	"Outsourcing/cmd/gorm/model"
	"fmt"

	"gorm.io/gorm"
)

const userTableName = "user"

var UserDao = new(userDao)

type userDao struct{}

// IUserDao user 表接口定义
type IUserDao interface {
	Insert(db *gorm.DB, user *model.User) (rowsAffected int64, err error)

	DeleteByUserID(db *gorm.DB, userId int) (rowsAffected int64, err error)

	Update(db *gorm.DB, user *model.User) (rowsAffected int64, err error)

	GetAllUser(db *gorm.DB) (userList []model.User, err error)

	LoadUserByUsername(db *gorm.DB, username string) (user *model.User, err error)
}

/*//
// Insert 添加用户
//	@receiver	u
//	@parameter	user 要添加到数据库中的用户对象
//	@return		rowsAffected 此次操作影响的行数
//	@return		err
//
func (userDao) Insert(db *gorm.DB, user *model.User) (rowsAffected int64, err error) {
	tx := db.Table(userTableName).Create(&user)
	if tx.Error != nil {
		return 0, err
	}
	return rowsAffected, nil
}

//
// DeleteByUserID 更具用户 ID 删除用户
//	@receiver	u
//	@parameter	userId 用户 ID
//	@return		rowsAffected 此次操作影响的行数
//	@return		err
//
func (userDao) DeleteByUserID(db *gorm.DB, userId int) (rowsAffected int64, err error) {
	// 根据主键删除
	tx := db.Table(userTableName).Delete(&model.User{}, userId)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

//
// Update 更新用户信息
//	@receiver	u
//	@parameter	user 包含要修改的用户信息
//	@return		rowsAffected 此次操作影响的行数, 更新成功, 如果更新了一个用户信息, 那么就返回 1, 如果更新了多个用户信息, 则返回被更新的用户数
//	@return		err
//
func (userDao) Update(db *gorm.DB, user *model.User) (rowsAffected int64, err error) {
	tx := db.Table(userTableName).Updates(user)
	if tx.Error != nil {
		fmt.Println("修改学生信息失败:", tx.Error)
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

//
// GetAllUser 查询所有用户信息
//	@receiver	u
//	@return		userList 用户信息列表
//	@return		err
//
func (userDao) GetAllUser(db *gorm.DB) (userList []model.User, err error) {
	userList = make([]model.User, 32)
	tx := db.Table(userTableName).Find(&userList)
	if tx.Error != nil {
		fmt.Println("查询全部学生信息失败:", tx.Error)
		return nil, tx.Error
	}

	return userList, nil
}*/

// LoadUserByUsername 根据用户名在数据库中查询用户
//
//	@receiver	userDao
//	@parameter	username 用户名
//	@return		user 查询到的用户对象, 如果查询期间发生错误, 则 user 为 nil
//	@return		err 错误信息
func (userDao) LoadUserByUsername(db *gorm.DB, username string) (*model.User, error) {
	var u model.User
	tx := db.Table(userTableName).Where("username=?", username).First(&u)
	if tx.Error != nil {
		fmt.Println("查询用户信息失败:", tx.Error)
		return nil, tx.Error
	}

	return &u, nil
}

/*//
// GetPagedUserList 分页查询所有的用户信息
//	@parameter	pageNum 页码(从 1 开始)
//	@parameter	pageSize 每页的记录数
//	@return		*orm.Page 分页对象
//
func (userDao) GetPagedUserList(db *gorm.DB, pageNum int, pageSize int) *orm.Page {
	var totalRows int64
	db.Table(userTableName).Count(&totalRows)

	userList := make([]model.User, 10)

	db.Table(userTableName).
		Scopes(orm.Paginate(pageNum, pageSize)).
		Find(&userList)

	return orm.NewPage(pageNum, pageSize, totalRows, userList)
}*/
