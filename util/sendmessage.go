package util

// 在指定qq群号里发送消息
func SendGroupMessage(groupID string, message string) error {
	// 发送消息
	_, _ = HttpGet("127.0.0.1:5700" + "/send_group_msg?group_id=" + groupID + "&message=" + message)
	return nil
}
