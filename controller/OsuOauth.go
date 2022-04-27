package controller

import (
	"KNBot/model"
	"KNBot/util"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"net/url"
)

/**
 * 关于Oauth第三方授权认证的一些参数。
 */
const (
	// 回调链接，后期更改
	redirect_uri = "http://localhost:5700/oauth2"
	// 申请的屙屎Oauth第三方应用程序给的id
	client_id = "14131"
	// 申请的屙屎Oauth第三方应用程序给的密钥
	client_secret = "HjMD0JFmHjeWdxspdF6f6H34RllMoikUZQgZn7Em"
)

/**
 * 拼接授权链接 * 此处的state日后再做处理
 * 会通过coolq来获取发送绑定命令者的一些参数
 * 处理这些参数作为id
 */
func AssembleAuthorizationUrl(state string) string {
	state = GetState()
	URL := "https://osu.ppy.sh/oauth/authorize" +
		"?state=" + state +
		"&redirect_uri=" + redirect_uri +
		"&scope=" + "friends.read identify public" +
		"&response_type=" + "code" +
		"&client_id=" + client_id
	return URL
}

/**
 * 初次获取到token
 * 需要根据返回的code去交换token
 * 初次根据token获取会同时返回access_token和refresh_token
 */
func GetToken(ctx iris.Context, code, state string) model.UserToken {
	// 交换token的url
	oauthUrl := "https://osu.ppy.sh/oauth/token"
	// 定义一个请求体
	body := make(url.Values)
	// 设置请求体参数
	body.Add("grant_type", "authorization_code")
	body.Add("code", code)
	body.Add("redirect_uri", redirect_uri)
	body.Add("client_id", client_id)
	body.Add("client_secret", client_secret)
	// 发送请求
	rsp := util.HttpPostForm(oauthUrl, body)

	var token model.UserToken
	// 读取响应
	dataByte := util.GetRspBody(rsp)
	// 将响应解析到token中
	err := json.Unmarshal(dataByte, &token)
	if err != nil {
		fmt.Println("解析失败")
	}

	return token
}

/**
 * 利用refresh_token刷新access_token
 */
func RefreshToken(ctx iris.Context, state string) model.UserToken {
	// 交换token的url
	oauthUrl := "https://osu.ppy.sh/oauth/token"
	// 定义一个请求体
	body := make(url.Values)
	// 设置请求体参数
	body.Add("client_id", client_id)
	body.Add("client_secret", client_secret)
	body.Add("refresh_token", OauthService.GetFreshToken(GetState()))
	body.Add("grant_type", "refresh_token")
	body.Add("redirect_uri", redirect_uri)
	// 发送请求
	rsp := util.HttpPostForm(oauthUrl, body)

	var token model.UserToken
	// 读取响应
	dataByte := util.GetRspBody(rsp)
	// 将响应解析到token中
	err := json.Unmarshal(dataByte, &token)
	if err != nil {
		fmt.Println("解析失败")
	}
	return token
}

/**
 * 机器人获取授权
 */
func GerBotAccessToken(ctx iris.Context, state string) model.BotToken {
	// 交换token的url
	oauthUrl := "https://osu.ppy.sh/oauth/token"
	// 定义一个请求体
	body := make(url.Values)
	// 设置请求体参数
	body.Add("grant_type", "client_credentials")
	body.Add("client_id", client_id)
	body.Add("client_secret", client_secret)
	body.Add("scope", "public")

	rsp := util.HttpPostForm(oauthUrl, body)

	var token model.BotToken
	// 读取响应
	dataByte := util.GetRspBody(rsp)
	// 将响应解析到token中
	err := json.Unmarshal(dataByte, &token)
	if err != nil {
		fmt.Println("解析失败")
	}
	return token
}
