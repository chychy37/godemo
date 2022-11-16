package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func ws(w http.ResponseWriter, r *http.Request) {
	var err error

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	for {
		// read blocked until new message received
		mt, p, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(mt)
		log.Println(string(p))

		// write
		err = c.WriteMessage(websocket.TextMessage, []byte("Received One Message"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", ws)
	http.ListenAndServe(":8888", nil)
}
