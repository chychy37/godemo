package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("Connecting to: 127.0.0.1:8888")

	c, ok := conn.(*net.UDPConn)
	if !ok {
		panic("Fail to convert Conn to UDPConn")
	}
	c.Write([]byte("hello"))
}
