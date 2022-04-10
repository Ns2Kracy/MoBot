package route

import (
	"KNBot/controller"
	"github.com/kataras/iris/v12"
)

func InitRoute(app *iris.Application) {

	app.Party("/")
	{
		Bind := app.Party("/bind")
		{
			Bind.Get("/", controller.AssembleAuthorizationUrl)
		}
	}
}
