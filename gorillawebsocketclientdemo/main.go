package main

import (
	"github.com/gorilla/websocket"
)

func main() {
	var err error

	// connect
	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8888/ws", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// write
	err = c.WriteMessage(websocket.TextMessage, []byte("hello"))
	if err != nil {
		panic(err)
	}
}
