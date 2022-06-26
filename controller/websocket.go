package controller

import (
	"MoBot/log"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"net/http"
	"sync"
	"time"
)

const (
	WriteWait      = 10 * time.Second
	PongWait       = 999999 * time.Second
	PingPeriod     = (PongWait * 9) / 10
	MaxMessageSize = 51200
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var maxConnId int64

type wsMessage struct {
	messageType int
	data        []byte
}

var WsConnAll map[int64]*WsConnection

type WsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *wsMessage // 读队列
	outChan  chan *wsMessage // 写队列

	mutex     sync.Mutex // 避免重复关闭管道,加锁处理
	isClosed  bool
	closeChan chan byte // 关闭通知
	id        int64
}

func WsHandler(ctx iris.Context) {
	// 应答客户端告知升级连接为websocket
	wsSocket, err := upgrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		log.GVA_LOG.Error("升级为websocket失败", zap.Any("err", err))
	}
	maxConnId++
	log.GVA_LOG.Info("新的连接", zap.Any("id", maxConnId))

	wsConn := &WsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
		id:        maxConnId,
	}

	WsConnAll[maxConnId] = wsConn
	// 处理器,发送定时信息，避免意外关闭
	go processLoop(wsConn)
	// 读协程panic: runtime error: invalid memory address or nil pointer dereference
	go wsReadLoop(wsConn)
	// 写协程
	go wsWriteLoop(wsConn)
}

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
func (wsConn *WsConnection) wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("连接已经关闭")
	}
	return nil
}

// 处理队列中的消息
func processLoop(wsConn *WsConnection) {
	// 处理消息队列中的消息
	// 获取到消息队列中的消息，处理完成后，发送消息给客户端
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			log.GVA_LOG.Error("获取消息出现错误:", zap.Any("err", err))
			wsConn.close()
			return
		}
		// log.Println(msg.messageType)
		// log.Println(string(msg.data))
		var msgData map[string]interface{}
		err = json.Unmarshal(msg.data, &msgData)
		if err != nil {
			log.GVA_LOG.Error("json信息解析错误", zap.Any("err", err))
		}
		log.GVA_LOG.Debug("收到消息：", zap.Any("msg", msgData))
		go HandleWsMsg(msgData)
	}
}

// 处理消息队列中的消息
func wsReadLoop(wsConn *WsConnection) {
	// 设置消息的最大长度
	wsConn.wsSocket.SetReadLimit(MaxMessageSize)
	err := wsConn.wsSocket.SetReadDeadline(time.Now().Add(PongWait))
	if err != nil {
		log.GVA_LOG.Error("wsSocket.SetReadDeadline failed", zap.Any("err", err))
	}
	for {
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			log.GVA_LOG.Error("消息读取出现错误", zap.Any("err", err))
			wsConn.close()
			// TODO 换成break和用return的区别？
			return
			//break
		}
		req := &wsMessage{
			msgType,
			data,
		}
		// 放入请求队列,消息入栈
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			return
			//break
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
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				log.GVA_LOG.Error("发送消息给客户端发生错误", zap.Any("err", err))
				// 切断服务
				wsConn.close()
				return
				//break
			}
		case <-wsConn.closeChan:
			// 获取到关闭通知
			return
			//break
		}
	}
}

// 关闭连接
func (wsConn *WsConnection) close() {
	log.GVA_LOG.Info("关闭连接")
	wsConn.wsSocket.Close()
	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if wsConn.isClosed == false {
		wsConn.isClosed = true
		// 删除这个连接的变量
		delete(WsConnAll, wsConn.id)
		close(wsConn.closeChan)
	}
}
