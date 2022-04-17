package provider

import (
	"KNBot/model"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
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
	state = "STATE"
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
func GetToken(ctx iris.Context, code, state string) model.Token {
	oauthUrl := "https://osu.ppy.sh/oauth/token"
	body := make(url.Values)
	body.Add("grant_type", "authorization_code")
	body.Add("code", code)
	body.Add("redirect_uri", redirect_uri)
	body.Add("client_id", client_id)
	body.Add("client_secret", client_secret)

	fmt.Println("截取到的code:", code)

	request, _ := http.PostForm(oauthUrl, body)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	defer request.Body.Close()
	//将响应json绑定到结构体
	var token model.Token
	dataByte, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(dataByte, &token)
	if err != nil {
		fmt.Println("解析失败")
	}

	fmt.Println("原始数据:", string(dataByte))
	return token
}

/**
 * 利用refresh_token刷新access_token
 */
func RefreshToken(ctx iris.Context, code, state string) model.Token {
	oauthUrl := "https://osu.ppy.sh/oauth/token"
	body := make(url.Values)
	body.Add("client_id", client_id)
	body.Add("client_secret", client_secret)
	body.Add("refresh_token", code)
	body.Add("grant_type", "refresh_token")
	body.Add("redirect_uri", redirect_uri)

	request, _ := http.PostForm(oauthUrl, body)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept-Type", "application/json")

	defer request.Body.Close()
	//将响应json绑定到结构体
	var token model.Token
	dataByte, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(dataByte, &token)
	if err != nil {
		fmt.Println("解析失败")
	}

	fmt.Println("原始数据:", string(dataByte))
	return token
}

/**
 * 机器人获取授权
 */
func GerBotAccessToken(ctx iris.Context, state string) model.Token {
	//获取code
	oauthUrl := "https://osu.ppy.sh/oauth/token"

	body := make(url.Values)
	body.Add("grant_type", "client_credentials")
	body.Add("redirect_uri", redirect_uri)
	body.Add("client_id", client_id)
	body.Add("client_secret", client_secret)
	body.Add("scope", "public")
	request, _ := http.PostForm(oauthUrl, body)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept-Type", "application/json")
	defer request.Body.Close()
	//将响应json绑定到结构体
	var token model.Token
	dataByte, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(dataByte, &token)
	if err != nil {
		fmt.Println("解析失败")
	}

	fmt.Println("原始数据:", string(dataByte))
	return token
}
