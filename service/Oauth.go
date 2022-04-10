package service

import (
	"KNBot/model"
	"xorm.io/xorm"
)

type OauthService struct {
	db *xorm.Engine
}

func (os *OauthService) GetClientId() int64 {
	var Oauth model.Oauth
	_, err := os.db.Where("client_id = ?").Get(&Oauth)
	if err != nil {
		return 0
	}
	return Oauth.ClientId
}

func (os *OauthService) GetClientSecret() string {
	var Oauth model.Oauth
	_, err := os.db.Where("client_secret = ?").Get(&Oauth)
	if err != nil {
		return ""
	}
	return Oauth.OauthToken
}

func (os *OauthService) GetRedirectUri() string {
	var Oauth model.Oauth
	_, err := os.db.Where("redirect_uri = ?").Get(&Oauth)
	if err != nil {
		return ""
	}
	return Oauth.RedirectUri
}
