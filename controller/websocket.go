package controller

import (
	"MoBot/log"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsMessage struct {
	msgType int
	msg     []byte
}

type WsConnection struct {
	wsConn      *websocket.Conn
	sendChan    chan *WsMessage
	receiveChan chan *WsMessage
	closeChan   chan byte
	mutex       sync.Mutex
	isClosed    bool
}

func (wsc *WsConnection) Send(msgType int, msg []byte) {
	wsc.sendChan <- &WsMessage{msgType: websocket.BinaryMessage, msg: msg}
}
func (wsc *WsConnection) Receive() *WsMessage {
	return <-wsc.receiveChan
}

func (wsc *WsConnection) Write() {
	for {
		select {
		// 取一个应答
		case msg := <-wsc.sendChan:
			// 写给websocket
			if err := wsc.wsConn.WriteMessage(websocket.BinaryMessage, msg.msg); err != nil {
				goto error
			}
		case <-wsc.closeChan:
			goto closed
		}
	}
error:
	wsc.Close()
closed:
}
func (wsc *WsConnection) Read() {
	for {
		// 读一个message
		_, msg, err := wsc.wsConn.ReadMessage()
		if err != nil {
			wsc.Close()
			return
		}
		req := &WsMessage{
			msgType: websocket.BinaryMessage,
			msg:     msg,
		}
		wsc.receiveChan <- req
	}

}

func (wsc *WsConnection) Process() {
	for {
		msg := wsc.Receive()
		if msg.msgType == websocket.CloseMessage {
			wsc.Close()
			break
		}
		var msgData map[string]interface{}
		err := json.Unmarshal(msg.msg, &msgData)
		if err != nil {
			log.GVA_LOG.Error("WsConnection Process error", zap.Error(err))
			continue
		}
		go HandleWsMsg(msgData)
	}
}

func (wsc *WsConnection) Close() {
	wsc.wsConn.Close()

	wsc.mutex.Lock()
	defer wsc.mutex.Unlock()
	if !wsc.isClosed {
		wsc.isClosed = true
		close(wsc.closeChan)
	}
}

func WsHandler(ctx iris.Context) {
	// 应答客户端告知升级连接为websocket
	wsSocket, err := upgrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		log.GVA_LOG.Error("WsHandler error", zap.Error(err))
		return
	}

	// 连接数保持一定数量，超过的部分不提供服务
	// 如果要控制连接数可以计算WsConnAll长度 len(WsConnAll)
	wsConn := &WsConnection{
		wsConn:      wsSocket,
		receiveChan: make(chan *WsMessage, 1000),
		sendChan:    make(chan *WsMessage, 1000),
		closeChan:   make(chan byte),
		isClosed:    false,
	}

	// 处理器,发送定时信息，避免意外关闭
	go wsConn.Process()
	// 读协程panic: runtime error: invalid memory address or nil pointer dereference

	go wsConn.Read()
	// 写协程
	go wsConn.Write()
}
