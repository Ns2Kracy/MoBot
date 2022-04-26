package route

import (
	"KNBot/controller"
	"github.com/kataras/iris/v12"
)

func InitRoute(app *iris.Application) {

	app.Party("/")
	{
		app.Get("/", controller.WsHandler)
		app.Get("/", controller.BindUrl)
		app.Get("/oauth2", controller.Oauth)
		Osu := app.Party("/osu")
		{
			Osu.Get("/setid", controller.SetOsuId)
			Info := Osu.Party("/info")
			{
				Info.Get("/me", controller.InfoMe)
			}
		}
	}
}
