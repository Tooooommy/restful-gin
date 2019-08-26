package connect

import (
	"CrownDaisy_GOGIN/im/config"
	"CrownDaisy_GOGIN/logger"
	"github.com/gorilla/websocket"
	"net/http"
)

type WebsocketHandler struct {
	Config      *config.WebSocketConfig
	Upgrader    *websocket.Upgrader
	pongHandler func(appData string) error
}

func New(c *config.WebSocketConfig) *WebsocketHandler {
	return &WebsocketHandler{
		Config: c,
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:    c.ReadBufferSize,
			WriteBufferSize:   c.WriteBufferSize,
			EnableCompression: c.Compression,
		},
	}
}

func (ws *WebsocketHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if ws.Config.CheckOrigin != nil {
		ws.Upgrader.CheckOrigin = ws.Config.CheckOrigin
	} else {
		ws.Upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	}
	conn, err := ws.Upgrader.Upgrade(rw, r, nil)
	if err != nil {
		logger.Sugar.Debugf("websocket upgrade error: %v", err)
		return
	}
	if ws.Config.Compression {
		err := conn.SetCompressionLevel(ws.Config.CompressionLevel)
		if err != nil {
			logger.Sugar.Debugf("websocket set compression level error:%v", err)
			return
		}
	}
	client := NewClient(conn)
	go client.Do()
}
