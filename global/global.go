package global

import (
	"MoBot/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

// 日志全局
var (
	GVA_DB     *xorm.Engine
	GVA_DBList map[string]*xorm.Engine
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	GVA_CONFIG config.MoBot
)

// 请求全局
const (
	Form_Type = "application/x-www-form-urlencoded"
	JSON_Type = "application/json"
)
