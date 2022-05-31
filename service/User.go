package service

import (
	"KNBot/model"
	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

type UserService struct {
	Db *xorm.Engine
}

/**
 * 保存验证过的用户信息
 */
func (us *UserService) SaveOauthUser(User model.User) error {
	//存取用户token
	_, err := us.Db.Insert(&User)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return nil
}

/**
 * 更新用户信息
 */
func (us *UserService) UpdateOauthUser(User model.User) error {
	_, err := us.Db.Where(" access_token = ? AND refresh_token = ? AND expire_in = ? ").Update(&User)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return nil
}
