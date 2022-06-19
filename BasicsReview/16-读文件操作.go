package main

import (
	"bufio"
	"fmt"
	"os"
)

func main16() {
	//如果文件不存在，创建文件
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()

	fmt.Println("文件打开成功")

	//创建一个带有缓冲区(用户缓冲)的reader
	reader := bufio.NewReader(f)
	//循环读取每一行数据
	for {
		//以换行符为结束符读取一行
		buf, err := reader.ReadBytes('\n')
		if err != nil && err.Error() == "EOF" { //读取到文件末尾
			fmt.Println("文件读取完毕") //自带换行符
			return
		} else if err != nil {
			fmt.Println("read file err:", err)
		}
		fmt.Print(string(buf))
	}
}
