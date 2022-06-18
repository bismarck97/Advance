package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//统计指定目录内单词出现次数：
//统计指定目录下，所有.txt文件中，“Love”这个单词 出现的次数。

//思路分析
//1.根据用户指定的目录，只读打开-->读目录的练习题。
//2.找到目录中的.txt，有可能有多个-->目录中找一 个指定类型文件
//3.打开其中一个.txt文件。循环读取一行 reader = bufio.NewReader, reader.ReadBytes("\n)-->读一行文件内容练习题
//4.将一行数据的字符串，拆分后,存入]string。Split, Fields-->字符串练习题
//5.遍历[string统计"Love" 单词出现的次数-->map练习题
func main() {
	fmt.Println("请输入您要打开的目录：")
	var path string
	fmt.Scan(&path)
	fmt.Println("请输入您要查找的字符串")
	var str string
	fmt.Scan(&str)
	//1.根据用户指定的目录，只读打开
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()
	//2.找到目录中的.txt，有可能有多个
	dir, err := f.ReadDir(-1)
	if err != nil {
		fmt.Println("read dir err:", err)
		return
	}
	var sum = 0
	for _, entry := range dir {
		//找出非文件类型
		if !entry.IsDir() {
			//找出包含.txt结尾的文件
			if strings.HasSuffix(entry.Name(), ".txt") {
				//创建函数来统计.txt中包含字符串的数量
				s := createStrSlice(path, entry.Name())
				//查找统计到的数据
				sum += findStr(s, str)
			}
		}
	}
	fmt.Printf("%s出现的次数是%d\n", str, sum)
}
func createStrSlice(path, fileName string) (s []string) {
	//打开传入的文件
	f, err := os.Open(path + "\\" + fileName)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()
	//3.打开其中一个.txt文件。循环读取一行
	reader := bufio.NewReader(f)
	//创建切片存入读取到的每一个字符串

	for {
		//读取每一行，用\n表示结尾
		bytes, err := reader.ReadBytes('\n')
		if err != nil && err.Error() == "EOF" {
			return
		} else if err != nil {
			fmt.Println("read bytes err:", bytes)
			return
		}
		//4.将一行数据的字符串，拆分后,存入[]string。
		strs := strings.Fields(string(bytes))
		for _, v := range strs {
			s = append(s, v)
		}
	}
}

//5.遍历[]string统计"Love" 单词出现的次数
func findStr(s []string, str string) int {
	var value int
	//创建map函数
	m := make(map[string]int)
	//遍历字符串切片,存入map中
	for _, s := range s {
		//判断key是否存在
		if _, ok := m[s]; ok {
			//存在则value+1
			m[s] += 1
		} else {
			//不存在则创建key
			m[s] = 1
		}
	}
	//根据对应的key找value
	for k, v := range m {
		if str == k {
			value += v
		}
	}
	return value
}
