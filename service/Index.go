package service

import (
	"xorm.io/xorm"
)

func NewOauthService(db *xorm.Engine) OauthService {
	return OauthService{
		db: db,
	}
}

func NewUserService(db *xorm.Engine) UserService {
	return UserService{
		db: db,
	}
}
