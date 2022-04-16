package controller

import (
	"KNBot/middleware"
	"KNBot/model"
	"fmt"
	"github.com/kataras/iris/v12"
	"strconv"
)

func BindUrl(ctx iris.Context) {
	ctx.Redirect(middleware.AssembleAuthorizationUrl("state"))
}

func Oauth(ctx iris.Context) {
	code := ctx.URLParam("code")
	state := ctx.URLParam("state")
	qq, _ := strconv.ParseInt(state, 10, 64)
	var token = middleware.GetToken(ctx, code, state)
	access_token := token.AccessToken
	refresh_token := token.RefreshToken

	fmt.Println("获取到的Access_Token:", token.AccessToken)
	fmt.Println("获取到的Refresh_Token", token.RefreshToken)

	expires_in := token.ExpiresIn
	User := model.User{
		Qq:           qq,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		ExpireTime:   int64(expires_in),
	}

	err := UserService.SaveOauthUser(User)
	if err != nil {
		fmt.Println("绑定出现错误")
		return
	}

	ctx.WriteString("成功绑定QQ:" + state)
}
