package model

import "time"

/**
 * 绑定的麦若群友信息
 */
type User struct {
	Id           int64
	Qq           int64
	OsuId        int64
	MainMode     int64
	AccessToken  string
	RefreshToken string
	Expiresin    int64
	JoinDate     time.Time
}

const (
	// 游戏模式
	osu = iota
	taiko
	fruits
	mania
)

const (
	ez = iota
	nf
	ht
	hr
	dt
	nc
	hd
	fl
)
