package controller

import (
	"KNBot/database"
	"KNBot/service"
)

var (
	db           = database.NewEngine()
	OauthService = service.NewOauthService(db)
	UserService  = service.NewUserService(db)
	OsuService   = service.NewOsuService(db)
)
