package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func ws(ws *websocket.Conn) {
	var err error

	defer ws.Close()

	for {
		// read blocked until new message received
		var msg string
		err = websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(msg)

		// write
		err = websocket.Message.Send(ws, "Received")
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(ws))
	http.ListenAndServe(":8888", nil)
}
