package main

import (
	"fmt"
	"net"
)

func main() {
	localAddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 1888}

	conn, err := net.ListenUDP("udp", localAddr)

	if err != nil {
		fmt.Printf("监听错误", err)
		return
	}
	defer conn.Close()

	data := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("监听错误", err)
			return
		}
		fmt.Printf("消息来源：%s,      消息内容：%s\n", addr, data[:n])

		remoteAddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 1889}
		_, err = conn.WriteToUDP([]byte("你好呀，这是我的回复"), remoteAddr)
		if err != nil {
			fmt.Printf("发送错误", err)
			return
		}

	}
}
