package datasource

import (
	"KNBot/config"
	"KNBot/model"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"xorm.io/xorm"
)

func NewEngine() *xorm.Engine {
	initConfig := config.InitConfig()
	if initConfig == nil {
		return nil
	}

	//database := initConfig.MySQLConfig

	//dataSourceName := database.User + ":" + database.Pwd + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.Database + "?charset=utf8"

	engine, err := xorm.NewEngine("mysql", "root:nk20021001@tcp(139.224.19.236:3306)/knbot?charset=utf8")

	if err != nil {
		panic(err.Error())
	}
	err = engine.Sync2(
		new(model.User),
	)
	err = engine.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	engine.ShowSQL(true)

	return engine
}