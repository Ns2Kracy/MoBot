package controller

import (
	"KNBot/model"
	"KNBot/provider"
	"fmt"
	"github.com/kataras/iris/v12"
)

func BindUrl(ctx iris.Context) {
	ctx.Redirect(provider.AssembleAuthorizationUrl("state"))
}

/**
 * 绑定用户
 * api: /oauth2
 */
func Oauth(ctx iris.Context) {
	// 获取code参数
	code := ctx.URLParam("code")
	// 获取state参数
	state := ctx.URLParam("state")
	// 得到access_token, refresh_token并绑定到用户
	var token = provider.GetToken(ctx, code, state)
	access_token := token.AccessToken
	refresh_token := token.RefreshToken
	expires_in := token.ExpiresIn

	// 新建一个 user 示例
	User := model.User{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		Expiresin:    int64(expires_in),
	}

	// 保存用户
	err := UserService.SaveOauthUser(User)
	if err != nil {
		fmt.Println("绑定出现错误")
		return
	}
}

func SetOsuId(ctx iris.Context) {

}
