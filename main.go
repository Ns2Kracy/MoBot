package main

import (
	"MoBot/config"
	"MoBot/database"
	"MoBot/log"
	"MoBot/router"
	"fmt"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

func main() {
	config.GVA_VP = config.Viper() //初始化Viper
	log.GVA_LOG = log.Zap()        //初始化zap日志库
	zap.ReplaceGlobals(log.GVA_LOG)
	config.GVA_DB = database.XormMysql() //初始化数据库
	if config.GVA_DB != nil {
		//创建数据库表
		database.RegisterTables(config.GVA_DB)
		//延迟关闭数据库
		db := config.GVA_DB.DB()
		defer db.Close()
	}

	//运行服务
	newApp()
}

func newApp() *iris.Application {
	app := iris.New()
	router.InitRoute(app)
	address := fmt.Sprintf(":%d", config.GVA_CONFIG.System.Port)
	if err := app.Run(iris.Addr(address), iris.WithoutServerError(iris.ErrServerClosed)); err != nil {
		log.GVA_LOG.Error("run server failed", zap.Error(err))
	}
	return app
}
