package main

import (
	"MoBot/core"
	"MoBot/database"
	"MoBot/global"
	"MoBot/router"
	"MoBot/util"
	"fmt"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = core.Viper() //初始化Viper
	global.GVA_LOG = util.Zap()  //初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = database.XormMysql() //初始化数据库
	if global.GVA_DB != nil {
		//创建数据库表
		database.RegisterTables(global.GVA_DB)
		//延迟关闭数据库
		db := global.GVA_DB.DB()
		defer db.Close()
	}

	//运行服务
	newApp()
}

func newApp() *iris.Application {
	app := iris.New()
	router.InitRoute(app)
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Port)
	if err := app.Run(iris.Addr(address), iris.WithoutServerError(iris.ErrServerClosed)); err != nil {
		global.GVA_LOG.Error("run server failed", zap.Error(err))
	}
	return app
}
