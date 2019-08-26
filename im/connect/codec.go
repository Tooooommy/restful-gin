package connect

import (
	"github.com/gorilla/websocket"
	"restful-gin/logger"
	"time"
)

type Message struct {
	Code    int    // 消息类型
	Content []byte // 消息体
}

type CodecAdapter interface {
	Decode(time.Duration) (*Message, error)
	Encode(*Message, time.Duration) error
}

// 编码器 :编码和解码
type BaseCodec struct {
	Conn *websocket.Conn
}

// 解码
func (c *BaseCodec) Decode(duration time.Duration) (*Message, error) {
	var msg = &Message{}
	_ = c.Conn.SetReadDeadline(time.Now().Add(duration))
	err := c.Conn.ReadJSON(msg)
	if err != nil {
		logger.Sugar.Debugf("read from conn error: %v", err)
		return nil, err
	}
	return msg, nil
}

// 编码
func (c *BaseCodec) Encode(msg *Message, duration time.Duration) error {
	_ = c.Conn.SetWriteDeadline(time.Now().Add(duration))
	err := c.Conn.WriteJSON(msg)
	if err != nil {
		logger.Sugar.Debugf("write to conn error: %v", err)
		return err
	}
	return nil
}
