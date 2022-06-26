package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	//获取用户键盘输入(stdin)，将输入数据发送给服务器
	go func() {
		b := make([]byte, 1024*4)
		for {
			n, err := os.Stdin.Read(b)
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				continue
			}
			//写给服务器 读多少，写多少
			conn.Write(b[:n])

		}
	}()
	buf := make([]byte, 1024*4)
	//回显服务器回发的大写数据
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("检测到服务器关闭，客户端也关闭")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("客户端读到服务器回发：", string(buf[:n]))
	}
}
