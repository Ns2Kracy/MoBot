package model

import "xorm.io/xorm"

//获取Oauth所需要的参数
type Oauth struct {
	ClientId     int64  `xorm:"not null pk autoincr INT(11)"`
	RedirectUri  string `xorm:"not null VARCHAR(255)"`
	ResponseType string `xorm:"not null VARCHAR(255)"`
	OauthToken   string `xorm:"not null VARCHAR(255)"`
}

var db *xorm.Engine

func GetClientId() (int64, error) {
	var oauth Oauth
	_, err := db.Get(&oauth)
	if err != nil {
		return 0, nil
	}
	return oauth.ClientId, nil
}

func GetRedirectUrl() (string, error) {
	var oauth Oauth
	_, err := db.Get(&oauth)
	if err != nil {
		return "", nil
	}
	return oauth.RedirectUri, nil
}

func GetOauthToken() (string, error) {
	var oauth Oauth
	_, err := db.Get(&oauth)
	if err != nil {
		return "", nil
	}
	return oauth.OauthToken, nil
}
