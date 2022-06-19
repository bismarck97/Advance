package main

import (
	"fmt"
	"io"
	"os"
)

func main15() {
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()

	fmt.Println("文件打开成功")

	n, err := f.WriteString("hello,world\r\n")
	if err != nil {
		fmt.Println("write string err:", err)
		return
	}
	fmt.Printf("写入 %d 个字节\n", n)

	i, _ := f.Seek(-5, io.SeekEnd)
	fmt.Printf("文件偏移 %d\n", i)

	n, _ = f.WriteAt([]byte("nihao"), i)
	fmt.Printf("写入 %d 个字节\n", n)
}
