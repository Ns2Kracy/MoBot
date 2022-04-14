package route

import (
	"KNBot/controller"
	"github.com/kataras/iris/v12"
)

func InitRoute(app *iris.Application) {

	app.Party("/")
	{
		Osu := app.Party("/")
		{
			Osu.Get("/", controller.BindUrl)
			Osu.Get("/oauth2", controller.Oauth)
		}
	}
}
