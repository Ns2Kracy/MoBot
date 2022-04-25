package controller

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
func GetToken(ctx iris.Context, code, state string) model.Token {
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
	request, _ := http.PostForm(oauthUrl, body)
	// 设定Header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	// 延迟关闭(十分重要)
	defer request.Body.Close()
	//将响应json绑定到结构体
	var token model.Token
	//读取响应
	dataByte, _ := ioutil.ReadAll(request.Body)
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
func RefreshToken(ctx iris.Context, state string) model.Token {
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
	request, _ := http.PostForm(oauthUrl, body)
	// 设定Header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept-Type", "application/json")
	// 延迟关闭(十分重要)
	defer request.Body.Close()
	// 将响应json绑定到结构体
	var token model.Token
	// 读取响应
	dataByte, _ := ioutil.ReadAll(request.Body)
	// 将响应解析到token中
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
	// 交换token的url
	oauthUrl := "https://osu.ppy.sh/oauth/token"
	// 定义一个请求体
	body := make(url.Values)
	// 设置请求体参数
	body.Add("grant_type", "client_credentials")
	body.Add("redirect_uri", redirect_uri)
	body.Add("client_id", client_id)
	body.Add("client_secret", client_secret)
	body.Add("scope", "public")
	// 发送请求
	request, _ := http.PostForm(oauthUrl, body)
	// 设定Header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept-Type", "application/json")
	// 延迟关闭(十分重要)
	defer request.Body.Close()
	//将响应json绑定到结构体
	var token model.Token
	// 读取响应
	dataByte, _ := ioutil.ReadAll(request.Body)
	// 将响应解析到token中
	err := json.Unmarshal(dataByte, &token)
	if err != nil {
		fmt.Println("解析失败")
	}

	fmt.Println("原始数据:", string(dataByte))
	return token
}

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
