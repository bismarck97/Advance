package main

import (
	"fmt"
	"os"
)

func main1401() {
	//创建文件，如果文件不存在则创建，如果文件存在覆盖原内容创建新的文件
	f, err := os.Create("./2.txt")
	if err != nil {
		fmt.Println("create err", err)
	}
	defer f.Close()
	fmt.Println("文件创建成功")
}
func main1402() {
	//只读方式打开文件
	f, err := os.Open("./2.txt")
	if err != nil {
		fmt.Println("open err", err)
	}
	defer f.Close()

	_, err = f.WriteString("hello world")
	if err != nil {
		fmt.Println("write err", err)
	} else {
		fmt.Println("文件写入成功")
	}

	fmt.Println("文件打开成功")
}
func main1403() {
	//以只读，只写，读写 方式打开文件
	f, err := os.OpenFile("./2.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("open err", err)
	}
	defer f.Close()

	_, err = f.WriteString("hello world")
	if err != nil {
		fmt.Println("write err", err)
	} else {
		fmt.Println("文件写入成功")
	}

	fmt.Println("文件打开成功")
}
