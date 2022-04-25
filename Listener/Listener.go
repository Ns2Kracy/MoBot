package Listener

func HandleWsMsg(msg map[string]interface{}) {
	switch msg["post_type"] {
	case "message":
		break
	default:
		break
	}
}
