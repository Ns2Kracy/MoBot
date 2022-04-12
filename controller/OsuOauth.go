package controller

import (
	"KNBot/model"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
)

const (
	redirect_uri  = "http://localhost:5700/oauth2"
	client_id     = "13964"
	client_secret = "QaMAxvh5APsw4HvEEziGRi0Ah8S06pM8wdztvS5B"
)

//授权链接
func AssembleAuthorizationUrl(state string) string {
	state = "2220496937"
	URL := "https://osu.ppy.sh/oauth/authorize" +
		"?state=" + state +
		"&redirect_uri=" + redirect_uri +
		"&scope=" + "friends.read identify public" +
		"&response_type=" + "code" +
		"&client_id=" + client_id
	return URL
}

//
//获取token的链接
func GetTokenUrl(code string) string {
	iris.New().Logger().Info("/token")
	URL := "https://osu.ppy.sh/oauth/token" +
		"?grant_type=" + "authorization_code" +
		"&code=" + code +
		"&redirect_uri=" + redirect_uri +
		"&client_id=" + client_id +
		"&client_secret=" + "QaMAxvh5APsw4HvEEziGRi0Ah8S06pM8wdztvS5B"
	return URL
}

func GetAccessToken(url string) (string, int64) {

	// 形成请求

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", 0
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	res, err = httpClient.Do(req)
	if err != nil {
		return "", 0
	}

	// 将响应体解析为 token，并返回
	var UsrToken model.Token
	if err = json.NewDecoder(res.Body).Decode(&UsrToken); err != nil {
		return "", 0
	}
	return UsrToken.AccessToken, int64(UsrToken.ExpiresIn)
}

func getRefreshToken(url string) (string, string, int64) {
	// 形成请求

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", "", 0
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	res, err = httpClient.Do(req)
	if err != nil {
		return "", "", 0
	}

	// 将响应体解析为 token，并返回
	var UsrToken model.Token
	if err = json.NewDecoder(res.Body).Decode(&UsrToken); err != nil {
		return "", "", 0
	}
	return UsrToken.AccessToken, UsrToken.RefreshToken, int64(UsrToken.ExpiresIn)
}

func BindUrl(ctx iris.Context) {

	//这里的state先就用state来测试一下，后面再改成用户的qq号
	Url := AssembleAuthorizationUrl("state")
	fmt.Println(Url)
	ctx.Redirect(Url, iris.StatusTemporaryRedirect)
}

//测试认证
func Oauth(ctx iris.Context) {
	// 获取 code
	var code = ctx.URLParam("code")
	fmt.Println("截取到的code:", code)

	// 通过 code, 获取 token
	var tokenAuthUrl = GetTokenUrl(code)
	fmt.Println("tokenAuthUrl:", tokenAuthUrl)

	var User model.User
	User.AccessToken, User.RefreshToken, User.ExpireTime = getRefreshToken(tokenAuthUrl)

	fmt.Println("获取的AccessToken:", User.AccessToken)
	fmt.Println("获取的RefreshToken:", User.RefreshToken)
	fmt.Println("获取的ExpireTime:", User.ExpireTime)
}
