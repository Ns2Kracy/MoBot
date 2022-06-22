package util

import (
	"fmt"
	"net/url"
)

// 在指定qq群号里发送消息
func SendGroupMessage(groupID string, message string) {
	fmt.Println("发送消息", "group_id="+groupID+"&message="+message)
	_ = NewHttpRequest("POST", "http://127.0.0.1:5700/send_group_msg", url.Values{"group_id": {groupID}, "message": {message}}, nil)
}

func SendPrivateMessage(userID string, message string) {
	fmt.Println("发送消息", "user_id="+userID+"&message="+message)
	_ = NewHttpRequest("POST", "http://127.0.0.1:5700/send_private_msg", url.Values{"user_id": {userID}, "message": {message}}, nil)
}
