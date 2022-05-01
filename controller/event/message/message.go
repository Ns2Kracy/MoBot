package message

func CheckType(msg map[string]interface{}) {
	switch msg["message_type"] {
	case "private":
		handlePrivate(msg)
		break
	case "group":
		handleGroup(msg)
		break
	default:
		break
	}
}

func handlePrivate(message map[string]interface{}) {

}

func handleGroup(message map[string]interface{}) {

}

// CheckPrivateMessage 检查私聊信息，返回对应的功能名称
func CheckPrivateMessage(msgStr string) (string, []string) {
	switch msgStr {
	case `/help`, `/帮助`:
		return "help", nil
	default:
		return "nil", nil
	}
}

// CheckPrivatemessage 检查群组信息，返回对应的功能名称
func CheckGroupMessage(msgStr string) (string, []string) {
	switch msgStr {
	case `/help`, `/帮助`:
		return "help", nil
	default:
		return "nil", nil
	}
}
