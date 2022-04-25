package controller

import (
	"github.com/kataras/iris/v12"
	"net/http"
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

/**
 * 通过绑定来获取个人信息，刷新osu_name以及osu_id
 */
func InfoMe(ctx iris.Context) {

	// 先检查是否绑定
	check := CheckToken(ctx)
	if check {
		osuUrl := baseUrl + "me" + "/" + "osu"

		client := &http.Client{}
		request, _ := http.NewRequest(http.MethodGet, osuUrl, nil)
		request.Header.Set("Authorization", "Bearer"+OauthService.GetAccessToken(GetState()))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept", "application/json")

		response, _ := client.Do(request)

		defer response.Body.Close()
		return
	} else {
		ctx.Redirect("/")
		return
	}
}
