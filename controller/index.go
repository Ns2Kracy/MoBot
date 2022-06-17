package controller

import (
	"MoBot/service"
)

var (
	OauthService = service.ServiceGroups.OsuService
	UserService  = service.ServiceGroups.UserService
	OsuService   = service.ServiceGroups.Oauth2Service
)
