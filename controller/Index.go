package controller

import (
	"KNBot/service"
	"xorm.io/xorm"
)

var (
	db           *xorm.Engine
	OauthService = service.NewOauthService(db)
)
