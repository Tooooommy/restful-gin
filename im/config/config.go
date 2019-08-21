package config

import "net/http"

type WebSocketConfig struct {
	Compression        bool
	CompressionLevel   int
	CompressionMinSize int
	ReadBufferSize     int
	WriteBufferSize    int
	CheckOrigin        func(r *http.Request) bool
}
