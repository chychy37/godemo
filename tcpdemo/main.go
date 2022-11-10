package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	// 由于这是一个tcp服务，listener贯穿整个程序的生命周期
	// 不需要主动执行Close()
	// defer listener.Close()

	for {
		// net包抽象了epoll模型，使用者只需要关注net.Conn对象
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	log.Println("Connecting to: ", c.RemoteAddr())

	// tcp为流数据，半包粘包问题的存在在于数据的边界
	// trick，使用bufio包来一次读取到特定字符
	r := bufio.NewReader(c) // bufsize=4096
	for {
		s, err := r.ReadString(byte('\n'))
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Receive data: " + s)
		c.Write([]byte(s))
	}
}
