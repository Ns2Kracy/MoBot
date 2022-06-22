package controller

import (
	"MoBot/config"
	"MoBot/log"
	"MoBot/plugins"
	"go.uber.org/zap"
	"strconv"
)

func HandleWsMsg(msg map[string]interface{}) {
	// fmt.Println(msg)
	switch msg["post_type"] {
	case "message":
		// 细分消息类型
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

// 将 global.GVA_CONFIG.System.Addr 转换为 ip:port
func GetAddrPort() string {
	return config.GVA_CONFIG.System.Host + strconv.Itoa(config.GVA_CONFIG.System.Port)
}

// 分发群消息
func HandleGroupMsg(msg map[string]interface{}) {

	log.GVA_LOG.Info("群聊消息", zap.Any("消息", msg))
	cmd := FilterMsg(msg)
	// 获取msg["group_id"].(float64)，并转换为int64
	groupId := int64(msg["group_id"].(float64))
	switch cmd {
	case `/setu`:
		break
	case `/ping`:
		SendGroupMessage(groupId, plugins.Ping())
		break

	}

}

// 分发私聊消息
func HandlePrivateMsg(msg map[string]interface{}) {

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

// 消息过滤, 仅上报以/开头的消息
func FilterMsg(msg map[string]interface{}) string {
	//将消息转换为字符串
	cmd, ok := msg["raw_message"].(string)
	if !ok {
		return ""

	}
	// 消息过滤, 仅上报以/开头的消息
	if cmd[0] != '/' {
		return ""

	}
	return cmd
}
