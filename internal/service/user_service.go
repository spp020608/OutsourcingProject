package service

import (
	"Outsourcing/cmd/gorm/model"
	"Outsourcing/cmd/gorm/query"
	"Outsourcing/internal/pkg/sms"
	"Outsourcing/internal/repository/cache"
	"Outsourcing/internal/repository/dao"
	"Outsourcing/model/dto"
	"Outsourcing/pkg/config"
	"Outsourcing/pkg/util"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func init() {
	cfg := config.Get().MySqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		zap.S().Info(err)
	}
	query.SetDefault(db)

}

var UserService = &userService{}

type userService struct{}

// SendCode
//
//	@Description:用于获取验证码
//	@receiver	userService
//	@parameter	identifyingCode
//	@return		result
func (userService) SendCode(identifyingCode *dto.IdentifyingCode) (result int) {
	code, ok := sms.SendCode(identifyingCode.Phone)
	if !ok {
		return 0
	}
	//同时将验证码存到缓存中
	cache.Set(identifyingCode.Phone, code, time.Minute*5)
	zap.S().Info("验证码为：", code)
	return 1
}

// Register 用户注册
//
//	@parameter	registerParam
//	@return		result 结果返回 不同情况
//	@return		err
func (userService) Register(registerParam *dto.RegisterParam) (result string, err error) {

	get, err := cache.Get(registerParam.Phone)
	if err != nil {
		zap.Error(err)
	}

	if registerParam.Code != get {
		zap.S().Info("验证码比对错误")
		return "验证码比对错误", nil
	}

	zap.S().Infof("用户注册接收到的密码为:%s\n", registerParam.Password)

	// 对密码进行加密
	zap.S().Infof("开始对明文密码进行加密, 明文密码为: %s\n", registerParam.Password)
	hashedPassword, err := util.GenerateFromPassword(registerParam.Password)
	zap.S().Infof("加密之后的密码为: %s\n", hashedPassword)
	if err != nil {
		zap.S().Errorf("密码加密失败: %s\n", err)
		return "密码加密失败", err
	}

	// 构造 User 对象
	user := &model.User{
		Username:    registerParam.Username,
		Password:    hashedPassword,
		Phone:       registerParam.Phone,
		IsPublisher: int32(registerParam.IdentityTag),
	}
	//开启协程的代码
	/*	q := query.Use(dao.DB)

		q.Transaction(func(tx *query.Query) error {
			if err := tx.User.Create(user); err != nil {
				return err
			}

			return nil
		})*/

	//比对数据库中的 username
	find, err := query.User.Where(query.User.Username.In(user.Username)).Find()
	if err != nil {
		zap.Error(err)
	}
	if len(find) != 0 {
		return "用户名已经存在", nil
	}

	//比对数据库中的 电话
	find, err = query.User.Where(query.User.Phone.In(user.Phone)).Find()
	if err != nil {
		zap.Error(err)
	}
	if len(find) != 0 {
		return "该电话号码已注册过该身份", nil
	}

	err = query.User.Create(user)
	if err != nil {
		zap.S().Error(err)
		return "数据库插入数据失败", err
	}

	return "注册用户成功", nil
}

// Login 用户登录
//
//	@parameter	loginInfo 前端传递的登录信息
//	@return		token 登录成功之后生成的 token
//	@return		err
func (userService) Login(loginInfo *dto.LoginInfo) (token string, err error) {
	zap.S().Infof("用户登录前端传递过来的密码为: %s", loginInfo.Password)
	// 判断传入参数是否为空
	if loginInfo == nil {
		zap.L().Info("登录信息为空 Login 直接返回")
		return "", errors.New("登录信息为空")
	}

	// 根据用户名从数据库中查询用户
	find, err := query.User.Where(query.User.Username.Eq(loginInfo.Username)).Find()
	if err != nil {
		zap.S().Error(err)
	}

	if err != nil {
		zap.S().Errorf("从数据库中查询用户名为: %s 的用户失败: %s", loginInfo.Username, err)
		return "", err
	}
	if len(find) == 0 {
		return "用户名不存在,请先注册", nil
	}

	user := find[0]

	zap.S().Infof("从数据库中获取的用户密码为:%s", user.Password)
	zap.S().Infof("接收到的前端用户用户密码为:%s", loginInfo.Password)

	// 比对数据库中的密码(加密之后的)和传递过来的密码(明文密码)
	err = util.CompareHashAndPassword(user.Password, loginInfo.Password)
	if err != nil {
		zap.S().Errorf("用户名或者密码错误: %s", err)
		return "", err
	}

	// 密码比对成功之后, 生成 token 字符串
	generateToken, err := util.GenerateToken(user.Username)
	if err != nil {
		zap.S().Errorf("根据用户名生成 token 失败: %s\n", err)
		return "", err
	}

	// 登录成功之后, 将登录用户信息保存到 redis 中
	if err = cache.SaveLoginUser(user); err != nil {
		zap.S().Infof("将登录用户存放到 redis 中失败: %s", err)
	}

	return generateToken, nil
}

// LoadUserByUsername
//
//	@Description:	用username查询 jwt验证用的
//	@receiver		userService
//	@parameter		username
//	@return			*model.User
//	@return			error
func (userService) LoadUserByUsername(username string) (*model.User, error) {
	// 首先判断 redis 中是否存在, 如果存在直接返回
	user, err := cache.GetLoginUser(username)
	if err != nil {
		zap.S().Errorf("从缓存中获取用户失败: %s", err)
		// 缓存中不存在, 去数据库中查找
		userInDB, err := dao.UserDao.LoadUserByUsername(dao.DB, username)
		if err != nil {
			zap.S().Errorf("从数据库中获取用户信息失败: %s", err)
			return nil, err
		}
		return userInDB, nil
	}
	return user, nil
}
