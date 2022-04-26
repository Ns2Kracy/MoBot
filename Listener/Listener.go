package Listener

func HandleWsMsg(msg map[string]interface{}) {
	switch msg["post_type"] {
	// case "message":
	case "message":
		switch msg["message_type"] {
		default:
			break
		}
		break
	// case "notice":
	case "notice":
		switch msg["notice_type"] {
		default:
			break
		}
		break
	// case "request":
	case "request":
		switch msg["request_type"] {
		default:
			break
		}
		break
	// case "meta_event":
	case "meta_event":
		switch msg["meta_event_type"] {
		default:
			break
		}
		break
	default:
		break
	}
}
