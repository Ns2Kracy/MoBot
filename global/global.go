package global

import (
	"MoBot/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

var (
	GVA_DB     *xorm.Engine
	GVA_DBList map[string]*xorm.Engine
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	GVA_CONFIG config.MoBot
)
