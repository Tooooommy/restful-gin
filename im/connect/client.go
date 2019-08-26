package connect

import (
	"github.com/gorilla/websocket"
	"restful-gin/logger"
	"time"
)

var (
	CodeSyncTrigger    = 1 // 消息同步
	CodeMessageSend    = 2 // 消息发送
	CodeMessageSendACK = 3 // 消息发送回执
	CodeMessage        = 4 // 消息投递
	CodeMessageACK     = 5 // 消息投递回执
	CodeHeartbeat      = 6 // 心跳
	CodeHeartbeatACK   = 7 // 心跳回执
)

type Client struct {
	Conn     *websocket.Conn
	Codec    CodecAdapter
	Broker   Broker
	IsSign   bool
	DeviceId int64
	UserId   int64
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Codec:  &BaseCodec{Conn: conn},
		Conn:   conn,
		Broker: NewEngine(),
	}
}

func (c *Client) SetCodec(codec CodecAdapter) {
	if codec != nil {
		c.Codec = codec
	}
}

func (c *Client) SetBroker(broker Broker) {
	if broker != nil {
		c.Broker = broker
	}
}

func (c *Client) Do() {
	err := c.Conn.SetReadDeadline(time.Now().Add(10 * time.Hour))
	if err != nil {
		logger.Sugar.Debugf("SetReadDeadline error:%v", err)
	}
	for {
		msg, err := c.Codec.Decode(10 * time.Hour)
		if err != nil {
			logger.Sugar.Debugf("message decode error:%v", err)
			return
		}
		c.Handle(msg)
	}
}

func (c *Client) Handle(msg *Message) {
	if c.IsSign == false {
		logger.Sugar.Infof("client not sign in")
		return
	}
	switch msg.Code {
	case CodeMessageSend:
		c.HandleMessageSend(msg)
	case CodeSyncTrigger:
		c.HandleSyncTrigger(msg)
	case CodeMessageSendACK:
		c.HandleMessageSendACK(msg)
	case CodeHeartbeat:
		c.HandleHeartbeat()
	}
}

func (c *Client) HandleMessageSend(msg *Message) {

}
func (c *Client) HandleMessageSendACK(msg *Message) {

}
func (c *Client) HandleSyncTrigger(msg *Message) {

}
func (c *Client) HandleHeartbeat() {
}

func (c *Client) Release() {
}

func (c *Client) SetPongHandler(pong func(appData string) error) {
	c.Conn.SetPongHandler(pong)
}

func (c *Client) SetPingHandler(ping func(appData string) error) {
	c.Conn.SetPingHandler(ping)
}

func (c *Client) SetClose(close func(code int, text string) error) {
	c.Conn.SetCloseHandler(close)
}

func (c *Client) SetHandleConnect(func()) {
}

func (c *Client) SetHandleDisconnect() {
}
