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
	_, err := us.db.Table("user").Where(" qq=? refresh_token=? access_token=? ", User.Qq, User.RefreshToken, User.AccessToken).Update(User)
	if err != nil {
		return false
	}
	return true
}

//更新用户信息
func (us *UserService) UpdateOauthUser(bindOsuId int64) error {
	var BindUser model.User
	_, err := us.db.ID(bindOsuId).Update(&BindUser)
	return err
}
