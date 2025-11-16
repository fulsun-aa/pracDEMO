package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	localAddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 1889}
	conn, err := net.ListenUDP("udp", localAddr)

	if err != nil {
		fmt.Print("监听打开失败", err)
		return
	}
	defer conn.Close()
	//远程地址
	remoteAddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 1888}
	_, err = conn.WriteToUDP([]byte("吉隆坡的天气总是翻云又赋鱼"), remoteAddr)

	if err != nil {
		fmt.Print("发送失败", err)
		return
	}
	resp := make([]byte, 1024)
	time.Sleep(time.Second)
	n, _, _ := conn.ReadFromUDP(resp)

	fmt.Println("收到回复：", string(resp[:n]))
}
