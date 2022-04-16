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
	ExpireTime   int64
	JoinDate     time.Time
}

const (
	// 游戏模式
	STD = iota
	TAIKO
	CTB
	MANIA
)
