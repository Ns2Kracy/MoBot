package config

import (
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

// 日志全局
var (
	GVA_DB     *xorm.Engine
	GVA_DBList map[string]*xorm.Engine
	GVA_VP     *viper.Viper
	GVA_CONFIG MoBot
)

// 请求全局
const (
	Form_Type = "application/x-www-form-urlencoded"
	JSON_Type = "application/json"
)
