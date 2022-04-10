package service

import (
	"xorm.io/xorm"
)

func NewOauthService(db *xorm.Engine) OauthService {
	return OauthService{
		db: db,
	}
}
