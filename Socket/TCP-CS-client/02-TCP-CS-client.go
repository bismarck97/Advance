package main

import (
	"fmt"
	"net"
)

func main02() {
	//指定服务器IP+port 创建 通信套接字
	conn, err := net.Dial("tcp", "192.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	//主动写数据给服务器
	conn.Write([]byte("hello"))
	//接收服务器回发的数据
	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("服务器回发:", string(buf[0:n]))
}
