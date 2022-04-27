package api

/*
 * 先把所有感觉会用的接口都写出来，然后再按照需求进行接口的组合
 */

var Host = "http://localhost:5700"

/**
 * 发送私聊消息
 * /send_private_msg
 * 参数说明：
 * user_id int64 对方的QQ号
 * message message 要发送的内容
 * auto_escape bool 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 message 字段是字符串时有效 , 默认为 false
 */
func SendPrivateMsg(user_id int64, group_id int64, message interface{}, auto_escape bool) {
}

/**
 * 发送群消息
 * /send_group_msg
 * 参数说明：
 * group_id string 群号
 * message message 要发送的内容
 * auto_escape bool 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 message 字段是字符串时有效 , 默认为 false
 */
func SendGroupMsg(group_id int64, message interface{}, auto_escape bool) {
}

/**
 * 发送消息
 * /send_msg
 * 参数说明：
 * message_type sting 消息类型, 支持 private、group , 分别对应私聊、群组, 如不传入, 则根据传入的 *_id 参数判断
 * user_id int64 对方的QQ号
 * group_id int64 群号
 * message message 要发送的内容
 * auto_escape bool 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 message 字段是字符串时有效 , 默认为 false
 */
func SendMsg(message_type string, user_id int64, group_id int64, message interface{}, auto_escape bool) {
}

/**
 * 撤回消息
 * /delete_msg
 * 参数说明：
 * message_id int32 消息ID
 */
func DeleteMsg(message_id int32) {
}

/**
 * 获取消息
 * /get_msg
 * 参数说明：
 * message_id int32 消息ID
 */
func GetMsg(message_id int32) {
}

/**
 * 群组踢人
 * /set_group_kick
 * 参数说明：
 * group_id int64 群号
 * user_id int64 要踢的qq号
 * reject_add_request bool 拒绝此人的加群请求 ， 默认为false
 */
func SetGroupKick(group_id int64, user_id int64, reject_add_request bool) {}

/**
 * 群组单人禁言
 * /set_group_ban
 * 参数说明：
 * group_id	int64	-	群号
 * user_id	int64	-	要踢的 QQ 号
 * reject_add_request	bool 拒绝此人的加群请求
 */
func SetGroupBan(group_id int64, user_id int64, reject_add_request bool) {

}

/**
 * 处理加好友请求
 * /set_friend_add_request
 * 参数说明：
 */
