package route

import (
	"KNBot/controller"
	"github.com/kataras/iris/v12"
)

func InitRoute(app *iris.Application) {

	app.Party("/")
	{
		Oauth := app.Party("/oauth")
		{
			Oauth.Get("/bind", controller.BindUser)
		}
	}
}
