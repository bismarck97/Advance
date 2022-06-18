package main

import (
	"fmt"
	"os"
)

//统计指定目录内单词出现次数：
//统计指定目录下，所有.txt文件中，“Love”这个单词 出现的次数。

//思路分析
//1.根据用户指定的目录，只读打开-->读目录的练习题。
//2.找到目录中的.txt，有可能有多个-->目录中找一 个指定类型文件
//3.打开其中-个.txt文件。循环读取一行 reader = bufio.NewReader, reader.ReadBytes("\n)-->读一行文件内容练习题
//4.将- -行数据的字符串，拆分后,存入]string。Split, Fields-->字符串练习题
//5.遍历[string统计"Love" 单词出现的次数-->map练习题
func main() {

	fmt.Println("请输入您要打开的目录：")
	var path string
	fmt.Scan(&path)
	//1.根据用户指定的目录，只读打开
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()

}
