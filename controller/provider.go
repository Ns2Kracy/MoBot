package controller

import "github.com/kataras/iris/v12"

/**
 * 一些常用的方法在这里写
 */

/**
 * 检查token是否存在
 */
func CheckToken(ctx iris.Context) bool {
	var state string
	access_token := OauthService.GetAccessToken(state)
	if access_token == "" {
		ctx.WriteString("该账号尚未绑定，请点击下面的链接进行绑定↓↓↓\n")
		return false
	}
	return true
}

/**
 * 通过coolq获取qq号用作state
 */
func GetState() string {
	var state string
	state = "state"
	return state
}
