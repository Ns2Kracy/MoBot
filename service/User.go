package service

import (
	"KNBot/model"
	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

type UserService struct {
	db *xorm.Engine
}

//保存验证过的用户信息
func (us *UserService) SaveOauthUser(User model.User) error {
	_, err := us.db.Table("user").
		Cols("access_token").
		Cols("refresh_token").
		Cols("qq").
		Cols("expire_in").
		Insert(&User)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return nil
}

//更新用户信息
func (us *UserService) UpdateOauthUser(User model.User) bool {
	return User.Id != 0
}
