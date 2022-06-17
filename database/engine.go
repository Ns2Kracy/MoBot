package database

import (
	"MoBot/global"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"os"
	"xorm.io/xorm"
)

//初始化xorm数据库
func XormMysql() *xorm.Engine {
	m := global.GVA_CONFIG.MySql
	//数据库名为空返回nil
	if m.Database == "" {
		return nil
	}
	//加载MySQL自定义配置
	dsn := m.Dsn()

	//初始化数据库
	db, err := xorm.NewEngine(m.Driver, dsn)
	if err != nil {
		global.GVA_LOG.Error("init mysql failed", zap.Error(err))
		os.Exit(0)
	}

	db.SetMaxIdleConns(m.MaxIdleConns)
	db.SetMaxOpenConns(m.MaxOpenConns)
	global.GVA_LOG.Info("init mysql success")
	return db
}

//if db, err := xorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
//	return nil
//} else {
//	sqlDB, _ := db.DB()
//	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
//	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
//	return db
//}

// 自动建立表结构
func RegisterTables(db *xorm.Engine) {
	err := db.Sync2()

	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
