package controller

import (
	"KNBot/model"
	"fmt"
	"github.com/kataras/iris/v12"
)

/**
 * 功能臆想
 * 1、获取自己的Osu个人信息
 * 2、获取其他人的Osu个人信息
 * 3、获取最近的游戏记录(不包含fail)
 * 4、获取最近的游戏记录(包含fail)
 * 5、获取今天的bps
 * 6、获取今天的tth
 * 7、获取今天的re
 * 8、今天打的最多的图是那张,打了多少次
 * 9、mania的能力评价
 * 10、std的能力评价
 * 11、查询自己bp列表上某一个bp
 */

var baseUrl = "http://localhost:5700"
var osuBaseUrl = "https://osu.ppy.sh/api/v2"

func BindUrl(ctx iris.Context) {
	ctx.WriteString(AssembleAuthorizationUrl("state"))
	ctx.Redirect(AssembleAuthorizationUrl("state"))
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
	var token = GetToken(ctx, code, state)
	access_token := token.AccessToken
	refresh_token := token.RefreshToken
	expires_in := token.ExpiresIn

	// 新建一个 user 示例
	User := model.User{
		OsuId:        GetId(),
		MainMode:     GetMode(),
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

func GetPlayInfo() {

}

/**
 * 设置主要模式，根据传入的state查询用户，再根据消息信息更新mode
 */
func SetMode(state string, mode int) {

}

func GetId() int   { return 0 }
func GetMode() int { return 0 }
