package model

/**
 * 绑定的麦若群友信息
 */
type User struct {
	Id           int64  `xorm:"pk autoincr"`
	Qq           int64  `xorm:"int index"`
	AccessToken  string `xorm:"varchar(255)"`
	RefreshToken string `xorm:"varchar(255)"`
	ExpireTime   int64  `xorm:"int(11)"`
}
