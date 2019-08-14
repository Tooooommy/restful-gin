package main

import (
	"CrownDaisy_GOGIN/logger"
	"golang.org/x/net/websocket"
	"io"
)

func EchoServer(ws *websocket.Conn) {
	_, _ = io.Copy(ws, ws)
}
func main() {
	//http.Handle("/echo", websocket.Handler(EchoServer))
	//err := http.ListenAndServe(":12345", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe " + err.Error())
	//}
	logger.Logger.Info("hi, man")
}
