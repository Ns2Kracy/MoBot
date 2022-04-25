package Bot

import (
	"KNBot/config"
)

// NewBot 返回一个Bot对象
func NewBot() *config.KnBotConfig {
	return &config.KnBotConfig{}
}

var PichuBot *config.KnBotConfig
