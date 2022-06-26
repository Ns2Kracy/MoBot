package main

import (
	"MoBot/config"
	"MoBot/controller"
	"MoBot/database"
	"MoBot/log"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"os"
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
	controller.WsConnAll = make(map[int64]*controller.WsConnection)
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Get("/", controller.WsHandler)
	if err := app.Run(iris.Addr(":"+config.GVA_CONFIG.System.WsPort), iris.WithoutServerError(iris.ErrServerClosed)); err != nil {
		log.GVA_LOG.Error("run server failed", zap.Error(err))
		os.Exit(1)
	}
}
