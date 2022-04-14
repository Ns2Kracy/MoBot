package service

import (
	"KNBot/model"
	"xorm.io/xorm"
)

type UserService struct {
	db *xorm.Engine
}

//保存验证过的用户信息
func (us *UserService) SaveOauthUser(User model.User) bool {
	return User.Id != 0
}

//更新用户信息
func (us *UserService) UpdateOauthUser(User model.User) bool {
	return User.Id != 0
}
