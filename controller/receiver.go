package controller

import "KNBot/controller/event/message"

func HandleMsg(msg map[string]interface{}) {
	switch msg["post_type"] {
	case "message":
		//消息事件
		message.CheckType(msg)
		break
	default:
		break
	}
}
