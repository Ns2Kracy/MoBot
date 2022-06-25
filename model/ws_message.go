package model

// 支持onebot标准的消息格式
type Message struct {
	PrivateParams
	GroupParams
	Action string `json:"action"`
	Echo   string `json:"echo"`
}

type PrivateParams struct {
	UserId  int64       `json:"user_id"`
	Message interface{} `json:"message"`
}
type GroupParams struct {
	GroupId int64       `json:"group_id"`
	Message interface{} `json:"message"`
}

type GroupMessage struct {
	Action string `json:"action"`
	Params struct {
		GroupId int64       `json:"group_id"`
		Message interface{} `json:"message"`
	}
}

type PrivateMessage struct {
	Action string `json:"action"`
	Params struct {
		UserId  int64       `json:"user_id"`
		Message interface{} `json:"message"`
	}
}

type CqMsg struct {
	Anonymous   interface{} `json:"anonymous"`
	Font        int         `json:"font"`
	GroupId     int         `json:"group_id"`
	Message     string      `json:"message"`
	MessageId   int         `json:"message_id"`
	MessageSeq  int         `json:"message_seq"`
	MessageType string      `json:"message_type"`
	PostType    string      `json:"post_type"`
	RawMessage  string      `json:"raw_message"`
	SelfId      int64       `json:"self_id"`
	Sender      struct {
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Card     string `json:"card"`
		Level    string `json:"level"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
		Sex      string `json:"sex"`
		Title    string `json:"title"`
		UserId   int64  `json:"user_id"`
	} `json:"sender"`
	SubType string `json:"sub_type"`
	Time    int    `json:"time"`
	UserId  int64  `json:"user_id"`
}
