package model

// 支持onebot标准的消息格式
type WsMessage struct {
	Action string `json:"action"`
	Params struct {
		GroupID string `json:"group_id"`
		Message string `json:"message"`
	} `json:"params"`
	Echo string `json:"echo"`
}
