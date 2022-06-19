package main

import (
	"fmt"
	"os"
)

func main18() {
	//获取用户输入的目录路径
	fmt.Println("请输入待查询的目录：")
	var path string
	fmt.Scan(&path)

	//打开目录
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()

	//读取目录项
	dir, err := f.ReadDir(-1) //读取目录中的所有目录项
	if err != nil {
		fmt.Println("read dir err:", err)
		return
	}
	//遍历返回的切片
	for _, entry := range dir {
		if entry.IsDir() { //是目录
			fmt.Printf("%s 是一个目录\n", entry.Name())
		} else {
			fmt.Printf("%s 是一个文件\n", entry.Name())
		}
	}
}
