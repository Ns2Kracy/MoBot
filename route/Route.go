package route

import (
	"KNBot/controller"
	"github.com/kataras/iris/v12"
)

func InitRoute(app *iris.Application) {

	app.Party("/")
	{
		app.Get("/", controller.BindUrl)
		app.Get("/oauth2", controller.Oauth)
		Osu := app.Party("/osu")
		{
			Osu.Get("/setid", controller.SetOsuId)
			Info := Osu.Party("/info")
			{
				Info.Get("/me", controller.InfoMe)
				Info.Get("/other", controller.InfoOther)
			}
		}
		Fun := app.Party("/fun")
		{
			Fun.Get("/setu")
		}
	}
}
