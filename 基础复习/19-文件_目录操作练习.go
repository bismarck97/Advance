package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//初级练习：
//指定目录检索特定文件：
//从用户给出的目录中，找出所有的 .flv 文件。
func main1901() {
	fmt.Println("请输入您要查找的目录")
	var path string
	fmt.Scan(&path)

	//打开文件目录
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()
	//读取目录项
	dir, err := f.ReadDir(-1) //读取所有目录
	if err != nil {
		fmt.Println("read dir err:", err)
		return
	}
	//遍历得到的文件切片
	for _, entry := range dir {
		if !entry.IsDir() { //判断是否是文件
			if strings.HasSuffix(entry.Name(), ".flv") {
				fmt.Println("包含flv的文件是:", entry.Name())
			}
		}
	}
}

//中级练习：
//指定目录拷贝特定文件：
//从用户给出的目录中，拷贝 .txt文件到指定目录中。
func main1902() {
	fmt.Println("请输入您要查找的目录")
	var path string
	fmt.Scan(&path)
	fmt.Println("请输入您要拷贝到的目录")
	var targetPath string
	fmt.Scan(&targetPath)
	//打开文件目录
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()
	//读取目录项
	dir, err := f.ReadDir(-1) //读取所有目录
	if err != nil {
		fmt.Println("read dir err:", err)
		return
	}
	//遍历得到的文件切片
	for _, entry := range dir {
		if !entry.IsDir() { //判断是否是文件
			if strings.HasSuffix(entry.Name(), ".txt") {
				copyDir(path, entry.Name(), targetPath)
			}
		}
	}
}

//高级练习：
//统计指定目录内单词出现次数：
//统计指定目录下，所有.txt文件中，“Love”这个单词 出现的次数。

func main() {

}

//拷贝文件到指定目录的操作
func copyDir(path, entryName, targetPath string) {
	//打开文件
	f1, err := os.Open(path + "\\" + entryName)
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer f1.Close()
	//写入文件 创建指定目录下文件
	f2, err := os.Create(targetPath + "\\" + entryName)
	if err != nil {
		fmt.Println("create err:", err)
		return
	}
	defer f2.Close()
	//从源文件中获取内容放入缓冲区
	buf := make([]byte, 4*1024)
	//循环从读文件中，获取数据"原封不动的"写到写文件中
	for {
		//读取数据写入到缓冲区切片中
		n, err := f1.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("文件读取完毕")
			return
		}
		//读多少，写多少
		f2.Write(buf[:n])
	}
}
