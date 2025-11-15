package main

import (
	"fmt"
	"net"
	"time"
)

func toCon() {
	conn, err := net.Dial("tcp", "127.0.0.1:1888")
	if err != nil {
		fmt.Print("请求打开失败", err)
		return
	}
	defer conn.Close()

	fmt.Println("连接成功", conn.RemoteAddr().String())
}

func chat() {
	conn, err := net.Dial("tcp", "127.0.0.1:1888")
	if err != nil {
		fmt.Print("请求失败", err)
		return
	}
	defer conn.Close()

	if err != nil {
		fmt.Print("连接失败", err)
		return
	}
	conn.Write([]byte("天空好像下雨"))
	time.Sleep(10 * time.Millisecond)

	var buf [1024]byte
	n, _ := conn.Read(buf[:])

	fmt.Print("收到回复：\n", string(buf[:n]))

}

func main() {
	chat()
}
