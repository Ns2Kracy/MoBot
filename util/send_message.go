package util

import (
	"fmt"
	"net/url"
)

// 在指定qq群号里发送消息
func SendGroupMessage(groupID string, message string) error {
	// 发送消息
	//_ = HttpGet("127.0.0.1:5700" + "/send_group_msg?group_id=" + groupID + "&message=" + message)
	fmt.Println("发送消息", "group_id="+groupID+"&message="+message)
	_ = NewHttpRequest("POST", "http://127.0.0.1:5700/send_group_msg", url.Values{"group_id": {groupID}, "message": {message}}, nil)
	/*	_ = HttpPostForm("/send_group_msg", url.Values{
		"group_id": {groupID}, "message": {message},
	})*/
	return nil
}
