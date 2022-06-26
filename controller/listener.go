package controller

func HandleWsMsg(msg map[string]interface{}) {
	// fmt.Println(msg)
	switch msg["post_type"] {
	case "message":
		switch msg["message_type"] {
		case "group":
			HandleGroupMsg(msg)
			break
		case "private":
			HandlePrivateMsg(msg)
			break
		default:
			break
		}
	case "notice":
		//通知事件
		HandleNoticeMsg(msg)
		break
	case "request":
		//请求事件
		HandleRequestMsg(msg)
		break
	case "meta_event":
		//元事件
		HandleMetaMsg(msg)
		break
	default:
		break
	}
}

// 分发群消息
func HandleGroupMsg(msg map[string]interface{}) {

	cmd := msg["raw_message"].(string)
	groupId := int64(msg["group_id"].(float64))
	switch cmd {
	case "/ping":
		SendGroupMessage(groupId, "Mooooooooooole")
		break
	default:
		break
	}
}

// 分发私聊消息
func HandlePrivateMsg(msg map[string]interface{}) {
	cmd := msg["raw_message"].(string)
	userId := int64(msg["user_id"].(float64))
	switch cmd {
	case "/ping":
		SendPrivateMessage(userId, "Mooooooooooole")
		break
	default:
		break
	}
}

// 分发通知消息
func HandleNoticeMsg(msg map[string]interface{}) {

}

// 分发请求消息
func HandleRequestMsg(msg map[string]interface{}) {

}

// 分发元事件
func HandleMetaMsg(msg map[string]interface{}) {

}
