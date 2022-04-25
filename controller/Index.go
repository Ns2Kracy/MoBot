package controller

import (
	"KNBot/datasource"
	"KNBot/service"
)

var (
	db           = datasource.NewEngine()
	OauthService = service.NewOauthService(db)
	UserService  = service.NewUserService(db)
	OsuService   = service.NewOsuService(db)
)
