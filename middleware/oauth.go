package middleware

import (
	"KNBot/model"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Oauth2中间件
const (
	redirect_uri  = "http://localhost:5700/oauth2"
	client_id     = "14131"
	client_secret = "HjMD0JFmHjeWdxspdF6f6H34RllMoikUZQgZn7Em"
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
func GetToken(ctx iris.Context, code, state string) model.Token {
	oauthUrl := "https://osu.ppy.sh/oauth/token"
	body := make(url.Values)
	body.Add("grant_type", "authorization_code")
	body.Add("code", code)
	body.Add("redirect_uri", redirect_uri)
	body.Add("client_id", client_id)
	body.Add("client_secret", client_secret)

	fmt.Println("截取到的code:", code)

	client := http.Client{}
	request, _ := http.NewRequest(http.MethodPost, oauthUrl, strings.NewReader(body.Encode()))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	req, _ := client.Do(request)

	defer req.Body.Close()
	//将响应json绑定到结构体
	var token model.Token
	dataByte, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(dataByte, &token)
	if err != nil {
		fmt.Println("解析失败")
	}

	fmt.Println("原始数据:", string(dataByte))
	return token
}
func RefreshToken(ctx iris.Context, code, state string) model.Token {
	oauthUrl := "https://osu.ppy.sh/oauth/token/"
	body := make(url.Values)
	body.Add("client_id", client_id)
	body.Add("client_secret", client_secret)
	body.Add("refresh_token", code)
	body.Add("grant_type", "refresh_token")
	body.Add("redirect_uri", redirect_uri)

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodPost, oauthUrl, strings.NewReader(body.Encode()))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept-Type", "application/json")
	req, _ := client.Do(request)

	defer req.Body.Close()

	//将响应json绑定到结构体
	var token model.Token
	dataByte, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(dataByte, &token)
	if err != nil {
		fmt.Println("解析失败")
	}

	return token
}
func GerBotAccessToken(ctx iris.Context, code, state string) {
	//获取code
	oauthUrl := "https://osu.ppy.sh/oauth/token"

	body := make(url.Values)
	body.Add("grant_type", "client_credentials")
	body.Add("code", code)
	body.Add("redirect_uri", redirect_uri)
	body.Add("client_id", client_id)
	body.Add("client_secret", client_secret)
	body.Add("scope", "public")
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodPost, oauthUrl, strings.NewReader(body.Encode()))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept-Type", "application/json")

	req, _ := client.Do(request)
	defer req.Body.Close()

	dataByte, _ := ioutil.ReadAll(req.Body)

	ctx.WriteString(string(dataByte))
}
