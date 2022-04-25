package config

import (
	"encoding/json"
	"os"
)

/**
 * 读取配置文件
 */
func InitBotConfig() *KnBotConfig {
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

func InitMysqlConfig() *MySQLConfig {
	var config *MySQLConfig
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
