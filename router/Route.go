package router

import "github.com/kataras/iris/v12"

func InitRoute(app *iris.Application) {

	app.Party("/")
	{
		app.Get("/")
	}
}
