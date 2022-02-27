package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"picture_community/entity/db"
	"picture_community/global"
	"picture_community/response"
	"picture_community/utils"
	"sync"
	"time"
)

// 用户类
type Client struct {
	FromId         uint // 发送者
	ToId           uint //接收者
	conn           *websocket.Conn
	MessageChannel chan string
}

// 用户管理
type ClientManager struct {
	Clients    map[uint]*Client
	Register   chan *Client
	Unregister chan *Client
}

var Manager = ClientManager{
	Clients:    make(map[uint]*Client),
	Register:   make(chan *Client, 1024),
	Unregister: make(chan *Client, 1024),
}

// 回复的消息
type ReplyMsg struct {
	StateCode int       `json:"state_code"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

var mux sync.RWMutex

func WsHandler(c *gin.Context) {

	up := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CheckOrigin解决跨域问题
			return true
		},
		Subprotocols: []string{c.Request.Header.Get("Sec-WebSocket-Protocol")},
	}

	uid, er := checkAuth(c, up.Subprotocols[0])
	if er != nil {
		response.UnAuthorization(c, nil, "请先登录")
		c.Abort()
		return
	}
	conn, err := up.Upgrade(c.Writer, c.Request, nil) // 升级成ws协议
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	// 创建一个用户实例
	client := &Client{
		FromId:         uid.(uint),
		ToId:           0,
		conn:           conn,
		MessageChannel: make(chan string),
	}
	// 用户注册到用户管理上
	Manager.Register <- client
	go client.Read()
	go client.Write()

}

func (c *Client) Write() {

	//pingTicker := time.NewTicker(time.Second * 120 )
	for {
		select {
		case message, ok := <-c.MessageChannel:
			if !ok {
				//_=c.conn.WriteMessage(websocket.CloseMessage,[]byte{})
				return
			}
			//log.Println(c.FromId,"接受消息:",string(message))
			replyMsg := ReplyMsg{
				StateCode: response.WebsocketOnlineReply,
				CreatedAt: time.Now(),
				Content:   message,
			}
			msg, _ := json.Marshal(replyMsg)

			_ = c.conn.WriteMessage(websocket.TextMessage, msg)
			//case <-pingTicker.C:
			//	// 服务端心跳:每120秒ping一次客户端，查看其是否在线
			//	_=c.conn.SetWriteDeadline(time.Now().Add(time.Second * 120))
			//	if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			//		return
			//	}

		}
	}
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c

	}()
	var receiveMsg map[string]string //ws的信息格式 json格式： {“contet”：“xxxxxx”}

	for {
		err := c.conn.ReadJSON(&receiveMsg) // 读取json格式，如果不是json格式，会报错
		if err != nil {
			//log.Println("数据格式不正确",err)
			return
		}

		chatMessage := db.ChatMessage{
			FromId:    c.FromId,
			ToId:      c.ToId,
			Content:   receiveMsg["content"],
			CreatedAt: time.Now(),
			IsRead:    0,
		}
		mux.RLock()
		cli, ok := Manager.Clients[c.ToId]
		mux.RUnlock()
		if ok {
			if c.FromId == cli.ToId {
				cli.MessageChannel <- receiveMsg["content"]
				chatMessage.IsRead = 1
			}
		}
		global.MysqlDB.Create(&chatMessage)

	}
}

func (manager *ClientManager) Start() {
	for {
		//log.Println("<---监听管道通信--->")
		select {
		case cli := <-Manager.Register: // 建立连接
			//log.Printf("建立新连接: %v", conn.ID)
			mux.Lock()
			Manager.Clients[cli.FromId] = cli
			mux.Unlock()
			replyMsg := &ReplyMsg{
				StateCode: response.WebsocketSuccess,
				CreatedAt: time.Now(),
				Content:   "已连接至服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = cli.conn.WriteMessage(websocket.TextMessage, msg)
		case cli := <-Manager.Unregister: // 断开连接

			_ = cli.conn.Close()
			close(cli.MessageChannel)
			mux.Lock()
			delete(Manager.Clients, cli.FromId)
			mux.Unlock()

		}
	}
}
func checkAuth(ctx *gin.Context, tokenString string) (interface{}, error) {
	if tokenString == "" {
		response.UnAuthorization(ctx, nil, "请先登录")
		//抛弃请求
		ctx.Abort()
		return -1, errors.New("请先登录")
	}
	claims, err := utils.ParserToken(tokenString)
	if err != nil {
		response.UnAuthorization(ctx, nil, err.Error())
		ctx.Abort()
		return -1, err
	}
	userId := claims.ID

	return uint(userId), nil
}
