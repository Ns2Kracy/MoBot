package controller

func HandleWsMsg(msg map[string]interface{}) {
	switch msg["type"].(string) {
	case "message":
	}
}
