package controller

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"log"
	"net/http"
	"sync"
)

const (
	MaxMessageSize = 51200
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type wsMessage struct {
	messageType int
	message     []byte
}

// 用于广播
var WsConnAll map[int64]*WsConnection

type WsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *wsMessage // 读队列
	outChan  chan *wsMessage // 写队列

	mutex     sync.Mutex // 避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知
}

var (
	RcvChan  chan *wsMessage
	SendChan chan *wsMessage
)

// 读取消息队列中的消息
func (wsConn *WsConnection) wsRead() (*wsMessage, error) {
	select {
	case msg := <-wsConn.inChan:
		// 获取到消息队列中的消息
		return msg, nil
	case <-wsConn.closeChan:
		return &wsMessage{}, errors.New("连接已经关闭")
	}
}

// 写入消息到队列中
func (wsConn *WsConnection) WsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("连接已经关闭")
	}
	return nil
}

// 处理队列中的消息
func processLoop(wsConn *WsConnection) {
	for {
		// 从队列中取出一个消息
		msg, err := wsConn.wsRead()
		if err != nil {
			log.Printf("wsRead error: %v", err)
		}
		var msgData map[string]interface{}
		err = json.Unmarshal(msg.message, &msgData)
		if err != nil {
			log.Printf("json.Unmarshal error: %v", err)
		}
		go HandleWsMsg(msgData)
	}
}

// 处理消息队列中的消息
func wsReadLoop(wsConn *WsConnection) {
	// 设置消息的最大长度
	wsConn.wsSocket.SetReadLimit(MaxMessageSize)
	for {
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			wsConn.wsClose()
			break
		}
		req := &wsMessage{
			msgType,
			data,
		}
		// 放入请求队列,消息入栈
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			break
		}
	}
}

// 发送消息给客户端
func wsWriteLoop(wsConn *WsConnection) {
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			// 写给websocket
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.message); err != nil {

				// 切断服务
				wsConn.wsClose()
				return
				//break
			}
		case <-wsConn.closeChan:
			// 关闭websocket
			wsConn.wsClose()
			break
		}
	}
}
func WsHandler(ctx iris.Context) {

	// 应答客户端告知升级连接为websocket
	wsSocket, err := wsUpgrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		return
	}
	wsConn := &WsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}

	// 处理器
	go processLoop(wsConn)
	// 读协程
	go wsReadLoop(wsConn)
	// 写协程
	go wsWriteLoop(wsConn)
}

func (wsConn *WsConnection) wsClose() {
	wsConn.wsSocket.Close()

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}
