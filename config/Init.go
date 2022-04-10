package config

import (
	"encoding/json"
	"os"
)

/**
 * 读取配置文件
 */
func InitConfig() *KnBotConfig {
	var config *KnBotConfig
	file, err := os.Open("./config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	return config
}
