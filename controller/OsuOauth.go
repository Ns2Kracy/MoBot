package controller

import (
	"KNBot/model"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

/**
 * 另外封装有关授权的比如
 * 拼接，因为需要操作到数据库，可能导致import循环
 * 所以先都放在这里
 * 春晚看看
 */

/**
 * 拼接授权链接
 * state qq
 * 此为现在osu设置的回调链接 http://localhost:5700/
 */

func AssembleAuthorizationUrl(state string) string {
	URL := "https://osu.ppy.sh/oauth/authorize" +
		"?state=" + state +
		"&redirect_uri=" + "http://localhost:5700/" +
		"&scope=" + "friends.read identify public" +
		"&response_type=" + "code" +
		"&client_id=" + strconv.FormatInt(OauthService.GetClientId(), 10)
	return URL
}

/**
 * 初次获取访问令牌,并返回json
 */
func GetAccessToken(User model.User) string {
	OsuUrl := "https://osu.ppy.sh/oauth/token"
	var body url.Values
	body.Add("client_id", strconv.FormatInt(OauthService.GetClientId(), 10))
	body.Add("client_secret", OauthService.GetClientSecret())
	body.Add("code", UserService.GetRefreshToken())
	body.Add("grant_type", "authorization_code")
	body.Add("redirect_uri", OauthService.GetRedirectUri())

	response, _ := http.PostForm(OsuUrl, body)
	response.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//获取返回的json
	bodyByte, _ := ioutil.ReadAll(response.Body)

	User.AccessToken = response.Header.Get("access_token")
	User.RefreshToken = response.Header.Get("refresh_token")
	User.ExpireTime, _ = strconv.ParseInt(response.Header.Get("expires_in"), 10, 64)

	UserService.SaveOauthUser(User)

	return string(bodyByte)
}

/**
 * 后续刷新访问令牌
 */

func RefreshToken(User model.User) {
	OsuUrl := "https://osu.ppy.sh/oauth/token"
	client := &http.Client{}
	var body url.Values
	body.Add("client_id", strconv.FormatInt(OauthService.GetClientId(), 10))
	body.Add("client_secret", OauthService.GetClientSecret())
	body.Add("code", UserService.GetRefreshToken())
	body.Add("grant_type", "refresh_token")
	body.Add("redirect_uri", OauthService.GetRedirectUri())

	response, _ := client.PostForm(OsuUrl, body)

	defer response.Body.Close()

	User.AccessToken = response.Header.Get("access_token")
	User.RefreshToken = response.Header.Get("refresh_token")
	User.ExpireTime, _ = strconv.ParseInt(response.Header.Get("expires_in"), 10, 64)

	UserService.UpdateOauthUser(User.OsuID)

}

func BindUser(ctx iris.Context) {
	//返回拼装url
	var User model.User
	Url := AssembleAuthorizationUrl("osu")
	ctx.Redirect(Url)
	//获取
	at := GetAccessToken(User)
	iris.New().Logger().Info(at)

	ctx.WriteString("成功绑定用户")
}
