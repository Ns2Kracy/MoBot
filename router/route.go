package router

import (
	"MoBot/controller"
	"github.com/kataras/iris/v12"
)

func InitRoute(app *iris.Application) {

	app.Party("/")
	{
		// 注册websocket
		app.Get("/", controller.WsHandler)

	}
}
