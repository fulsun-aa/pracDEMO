package main

import (
	"fmt"
	"log"
	"net"
)

func closeWait(conn net.Conn) {
	//conn.Close()

	data := make([]byte, 1024)

	_, err := conn.Read(data)

	if err != nil {
		log.Fatal("读取错误", err)
		return
	}

	fmt.Println("收到消息：", string(data))

}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8010")

	if err != nil {
		log.Fatal("读取错误", err)

	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("读取错误", err)
			return

		}
		go closeWait(conn)
	}

}
