package database

import (
	_ "github.com/go-sql-driver/mysql"
	"log"

	"xorm.io/xorm"
)

func NewEngine() *xorm.Engine {
	/*
		initConfig := config.InitMysqlConfig()
		if initConfig == nil {
			return nil
		}
	*/

	//database := initConfig

	//dataSourceName := database.User + ":" + database.Pwd + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.Database + "?charset=utf8"
	//root:nk20021001@tcp(139.224.19.236:3306)/knbot?charset=utf8
	engine, err := xorm.NewEngine("mysql", "root:nk20021001@tcp(139.224.19.236:3306)/knbot?charset=utf8")

	if err != nil {
		panic(err.Error())
	}
	err = engine.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	engine.ShowSQL(true)

	return engine
}
