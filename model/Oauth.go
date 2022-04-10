package model

//获取Oauth所需要的参数
type Oauth struct {
	ClientId     int64  `xorm:"not null pk autoincr INT(11)"`
	RedirectUri  string `xorm:"not null VARCHAR(255)"`
	ResponseType string `xorm:"not null VARCHAR(255)"`
	Scope        string `xorm:"not null VARCHAR(255)"`
	OauthId      int64  `xorm:"not null pk autoincr INT(11)"`
	OauthToken   string `xorm:"not null VARCHAR(255)"`
	Url          string `xorm:"not null VARCHAR(255)"`
}
