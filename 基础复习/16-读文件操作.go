package main

import (
	"fmt"
	"os"
)

func main() {
	//如果文件不存在，创建文件
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()

	fmt.Println("文件打开成功")
}
