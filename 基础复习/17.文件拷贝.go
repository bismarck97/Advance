package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//打开读文件
	f1, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f1.Close()
	//创建写文件
	f2, err := os.OpenFile("test2.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f2.Close()
	//从读文件中获取数据，放到缓冲区中
	buf := make([]byte, 4*1024)
	//循环从读文件中，获取数据"原封不动的"写到写文件中
	for {
		n, err := f1.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Printf("读取完毕 n = %d\n", n)
			return
		}
		f2.Write(buf[:n]) //读多少，写多少
	}

}
