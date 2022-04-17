package service

import (
	"KNBot/model"
	"xorm.io/xorm"
)

type OauthService struct {
	db *xorm.Engine
}

func (os *OauthService) GetAccessToken(User model.User) string {
	err := os.db.Table("user").Where(" qq = ?").Find(&User)
	if err != nil {
		return "token失效或者无此token"
	}
	return User.AccessToken
}
