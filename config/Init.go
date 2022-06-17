package config

import "github.com/spf13/viper"

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.AddConfigPath(".")
	Config.SetConfigType("yaml")

	err := Config.ReadInConfig()

	if err != nil {
		panic(err)
	}
}
