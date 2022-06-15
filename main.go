package main

import (
	"MoBot/router"
	"github.com/kataras/iris/v12"
)

func main() {
	app := newApp()
	app.Run(iris.Addr(":5700"))
}

func newApp() *iris.Application {
	app := iris.New()
	router.InitRoute(app)
	app.Logger().SetLevel("debug")

	return app
}
