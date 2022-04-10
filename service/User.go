package service

import (
	"KNBot/model"
	"xorm.io/xorm"
)

type UserService struct {
	db *xorm.Engine
}

//保存验证过的用户信息
func (us *UserService) SaveOauthUser(BindUser model.User) error {
	_, err := us.db.ID(BindUser.Id).Insert(&BindUser)
	return err
}

//更新用户信息
func (us *UserService) UpdateOauthUser(bindOsuId int64) error {
	var BindUser model.User
	_, err := us.db.ID(bindOsuId).Update(&BindUser)
	return err
}

func (us *UserService) GetRefreshToken() string {
	var user model.User
	_, err := us.db.Where(" id = ?").Get(&user)
	if err != nil {
		return ""
	}
	return user.RefreshToken
}
