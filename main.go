package main

import (
	"KNBot/router"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	router.InitRoute(app)
	app.Logger().SetLevel("debug")
	app.Run(iris.Addr(":5700"))
}
