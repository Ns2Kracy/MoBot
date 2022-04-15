package controller

import (
	"KNBot/model"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

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
	oauthUrl := "https://osu.ppy.sh/oauth/token/"
	body := make(url.Values)
	body.Add("grant_type", "authorization_code")
	body.Add("code", code)
	body.Add("redirect_uri", redirect_uri)
	body.Add("client_id", client_id)
	body.Add("client_secret", client_secret)

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
	fmt.Println("原始json数据:", string(dataByte))

	fmt.Println("获取到的Access_Token:", token.AccessToken)
	fmt.Println("获取到的Refresh_Token", token.RefreshToken)

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
	request.Header.Set("Accept-Type", "application/x-www-form-urlencoded")
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
	request.Header.Set("Accept-Type", "application/x-www-form-urlencoded")

	req, _ := client.Do(request)
	defer req.Body.Close()

	dataByte, _ := ioutil.ReadAll(req.Body)

	ctx.WriteString(string(dataByte))
}

func BindUrl(ctx iris.Context) {
	ctx.Redirect(AssembleAuthorizationUrl("2220496937"))
}

func Oauth(ctx iris.Context) {
	code := ctx.URLParam("code")
	state := ctx.URLParam("state")
	qq, _ := strconv.ParseInt(state, 10, 64)
	var token = GetToken(ctx, code, state)
	access_token := token.AccessToken
	refresh_token := token.RefreshToken
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
