package main

import (
	"KNBot/route"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	route.InitRoute(app)
	app.Logger().SetLevel("debug")
	app.Listen(":5700")
}
