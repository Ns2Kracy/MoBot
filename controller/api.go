package controller

import (
	"MoBot/log"
	"MoBot/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var wsEvent *websocket.Conn

func initWsEvent(conn *websocket.Conn) {
	wsEvent = conn
}

// SendPrivateMessage 发送私聊消息
func SendPrivateMessage(UserId int64, message interface{}) {
	msg := model.PrivateMessage{
		Action: "/send_private_msg",
		Params: struct {
			UserId  int64       `json:"user_id"`
			Message interface{} `json:"message"`
		}{UserId: UserId, Message: message},
	}
	msgData, err := json.Marshal(msg)
	if err != nil {
		log.GVA_LOG.Error("SendprivateMessage error", zap.Error(err))
	}
	fmt.Println(string(msgData))
	// 调用websocket发送消息
	var WsApi WsConnection
	WsApi.Send(websocket.BinaryMessage, msgData)
}

// SendGroupMessage 发送群聊消息
func SendGroupMessage(GroupId int64, message interface{}) {
	msg := model.GroupMessage{
		Action: "",
		Params: struct {
			GroupId int64       `json:"group_id"`
			Message interface{} `json:"message"`
		}{GroupId: GroupId, Message: message},
	}

	msgData, err := json.Marshal(msg)
	if err != nil {
		log.GVA_LOG.Error("SendGroupMessage error", zap.Error(err))
	}
	var WsApi WsConnection
	WsApi.Send(websocket.BinaryMessage, msgData)
}
