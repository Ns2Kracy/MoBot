package util

import "net/url"

var host = "http://localhost:5700"

func SendPrivate(qq string, msg string) {
	_, _ = HttpGet(host + "/send_private_msg?user_id=" + qq + "&message=" + msg)
}

func SendGroup(group string, msg string) {
	_, _ = HttpGet(host + "/send_group_msg?group_id=" + group + "&message=" + msg)
}

func SendGroupPost(group string, msg string) {
	var data url.Values
	data.Set("group_id", group)
	data.Set("message", msg)
	_, _ = HttpPostForm(host+"/send_group_msg", data)
}
