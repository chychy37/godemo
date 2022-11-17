package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8888})
	if err != nil {
		panic(err)
	}
	// 这里同tcpdemo，不需要手动调用Close()
	// defer conn.Close()

	for {
		// MTU = 1500
		// UDP max data len = 65535-8-20
		// 如果upd包在ip层被切割，一个下层包丢失将导致整个udp包丢失
		// UDP recommended data len = 1500-8-20 （需要双方约定好）
		data := make([]byte, 1472)
		// UDP 接包顺序是不确定的，也可能永远接收不到
		_, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(addr.String() + " Receive data: " + string(data))
	}
}
