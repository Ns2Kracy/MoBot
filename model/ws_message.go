package model

// 支持onebot标准的消息格式
type Message struct {
	Params
	Action string `json:"action"`
	Echo   string `json:"echo"`
}

type Params struct {
	UserId  string `json:"user_id"`
	GroupID string `json:"group_id"`
	Message string `json:"message"`
}
