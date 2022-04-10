package model

/**
 * 绑定的麦若群友信息
 */
type User struct {
	Id           int64  `xorm:"pk autoincr"`
	Qq           int64  `xorm:"int index"`
	OsuName      string `xorm:"varchar(64) index"`
	OsuID        int64  `xorm:"int(11) notnull unique"`
	AccessToken  string `xorm:"varchar(255)"`
	RefreshToken string `xorm:"varchar(255)"`
	ExpireTime   int64  `xorm:"int(11)"`
}
