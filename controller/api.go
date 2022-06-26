package controller

import (
	"MoBot/config"
	"MoBot/util"
	"strconv"
)

// SendMessage 发送消息
func SendMessage(UserId int64, message interface{}) {

}

func SendPrivateMessage(userId int64, msg string) {
	_ = util.HttpGet(config.Http_Url + "/send_private_msg?user_id=" + strconv.FormatInt(userId, 10) + "&message=" + msg)
}

// SendGroupMessage 发送群聊消息
func SendGroupMessage(GroupId int64, msg string) {
	_ = util.HttpGet(config.Http_Url + "/send_group_msg?group_id=" + strconv.FormatInt(GroupId, 10) + "&message=" + msg)
}
