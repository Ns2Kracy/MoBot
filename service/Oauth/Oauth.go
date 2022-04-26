package Oauth

import (
	"KNBot/model"
	"xorm.io/xorm"
)

type Oauth2Service struct {
	Db *xorm.Engine
}

func (os *Oauth2Service) GetAccessToken(state string) string {
	var User model.User
	exist, _ := os.Db.Table("user").Where(" qq = ? ", state).Exist("access_token")
	if !exist {
		return "token失效或者无此token"
	}
	return User.AccessToken
}

func (os *Oauth2Service) GetFreshToken(state string) string {
	var User model.User
	exist, _ := os.Db.Table("user").Where(" qq = ? ", state).Get(&User)
	if !exist {
		return "token失效或者无此token"
	}
	return User.AccessToken
}

func (os *Oauth2Service) GetExpiresIn(state string) int64 {
	return 1
}
