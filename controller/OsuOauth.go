package controller

import "github.com/kataras/iris/v12"

/**
 * 拼接授权链接,这是绑定步骤的第一步
 * state qq
 * 此为现在osu设置的回调链接 http://localhost:5700/
 */
func AssembleAuthorizationUrl(ctx iris.Context) {
	// 获取发送用户的qq号
	qq := "2220496937"
	// 获取发送用户的群号
	group := "928936255"
	// 得到state
	state := qq + "+" + group
	// 获取redirect_uri
	redirect_uri := "http://localhost:5700/"
	// 获取scope
	scope := "friends.read identify public"
	// 拼接授权链接
	url := "https://osu.ppy.sh/oauth/authorize" +
		"?state=" + state +
		"&redirect_uri=" + redirect_uri +
		"&scope=" + scope +
		"&response_type=" + "code" +
		"&client_id=" + "21269776"
	// 返回授权链接
	ctx.WriteString("请在此链接完成绑定->" + url)
}

/**
 * 获取访问令牌,将用户的oauth存入数据库,这是第二步
 */
func GetAccessToken(ctx iris.Context) {
}

/**
 * 刷新访问令牌
 */
