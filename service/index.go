package service

type ServiceGroup struct {
	UserService
	OsuService
	Oauth2Service
}

var ServiceGroups = new(ServiceGroup)
