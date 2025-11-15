package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func waitCon() {
	listener, err := net.Listen("tcp", "127.0.0.1:1888")
	if err != nil {
		fmt.Print("监听打开失败", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("连接失败", err)
			return
		}
		fmt.Print("连接成功：来自", conn.RemoteAddr().String())
	}

}

func chat() {
	listen, err := net.Listen("tcp", "127.0.0.1:1888")
	if err != nil {
		fmt.Print("监听打开失败", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Print("链接打开失败", err)
			return
		}
		go readWrite(conn)

	}
}
func readWrite(con net.Conn) {
	defer con.Close()

	reader := bufio.NewReader(con)

	for {
		var buf [1024]byte
		n, err := reader.Read(buf[:])

		if err != nil && err != io.EOF {
			fmt.Print("读取失败", err)
			break
		}
		respond := string(buf[:n])
		fmt.Println(respond)

		con.Write([]byte("我好想住你隔壁"))
	}

}

func main() {
	chat()
}
