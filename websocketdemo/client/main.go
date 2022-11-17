package main

import (
	"golang.org/x/net/websocket"
)

func main() {
	var err error

	// connect
	ws, err := websocket.Dial("ws://127.0.0.1:8888/ws", "", "http://127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	// write
	err = websocket.Message.Send(ws, "hello")
	if err != nil {
		panic(err)
	}
}
