package service

import (
	"xorm.io/xorm"
)

func NewOauthService(db *xorm.Engine) Oauth2Service {
	return Oauth2Service{
		Db: db,
	}
}

func NewUserService(db *xorm.Engine) UserService {
	return UserService{
		Db: db,
	}
}

func NewOsuService(db *xorm.Engine) OsuService {
	return OsuService{
		Db: db,
	}
}
