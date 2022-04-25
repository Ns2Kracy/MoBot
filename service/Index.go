package service

import (
	"KNBot/service/Oauth"
	"KNBot/service/OsuUser"
	"KNBot/service/User"
	"xorm.io/xorm"
)

func NewOauthService(db *xorm.Engine) Oauth.Oauth2Service {
	return Oauth.Oauth2Service{
		Db: db,
	}
}

func NewUserService(db *xorm.Engine) User.UserService {
	return User.UserService{
		Db: db,
	}
}

func NewOsuService(db *xorm.Engine) OsuUser.OsuService {
	return OsuUser.OsuService{
		Db: db,
	}
}
