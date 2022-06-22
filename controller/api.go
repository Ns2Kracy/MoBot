package controller

import (
	"MoBot/config"
	"MoBot/model"
	"MoBot/util"
	"encoding/json"
	"github.com/gorilla/websocket"
	"strconv"
)

func SendWsMessage(msg string) {
	// 调用WsWrite
	wsConn := WsConnection{}
	wsConn.WsWrite(websocket.TextMessage, []byte(msg))
}

func SendPrivateMessage(UserId, GroupId int64, msg string) {
	Message := &model.Message{
		Params: model.Params{
			UserId:  strconv.FormatInt(UserId, 10),
			GroupID: strconv.FormatInt(GroupId, 10),
			Message: msg,
		},
		Action: "/send_private_msg",
		Echo:   "",
	}
	// 将Message转换为json格式
	message, _ := json.Marshal(Message)
	_ = util.NewHttpRequest("POST", "127.0.0.1:5700/", message, config.Form_Type)
}

func SendGroupMessage(GroupId int64, msg string) {
	Message := &model.Message{
		Params: model.Params{
			GroupID: strconv.FormatInt(GroupId, 10),
			Message: msg,
		},
		Action: "/send_group_msg",
		Echo:   "",
	}
	// 将Message转换为json格式
	message, _ := json.Marshal(Message)

	_ = util.NewHttpRequest("POST", "127.0.0.1:5700/", message, config.Form_Type)
}
