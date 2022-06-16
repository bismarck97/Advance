package main

import (
	"fmt"
)

//练习1：给定一个字符串列表，在原有slice上返回不包含空字符串的列表， 如：
//{"red", "", "black", "", "", "pink", "blue"}
//——> {"red", "black", "pink", "blue"}
//练习2：写一个函数，就地消除[]string中重复字符串，如：
//{"red", "black", "red", "pink", "blue", "pink", "blue"}
//——>	{"red", "black", "pink", "blue"}

func main0301() {
	data := []string{"red", "", "black", "", "", "pink", "blue"}
	afterData := noEmpty2(data)
	fmt.Println(afterData)
}
func main0302() {
	data := []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	afterData := noRepetition(data)
	fmt.Println(afterData)
}

func noRepetition(data []string) []string {
	//out := make([]string, 1)
	//out[0] = data[0]
	out := data[:1]
	//遍历原始切片字符串
	for _, str := range data {
		//设置flag标志位
		i := 0
		//比较取出的str是否存在
		for ; i < len(out); i++ {
			if str == out[i] {
				//如果重复就跳出本次循环
				break
			}
		}
		//判断循环次数是否等于切片长度，如果等于说明循环正常结束，没有经过break跳过，可以把数据值填入
		if i == len(out) {
			out = append(out, str)
		}
	}
	return out
}

///练习：给定一个字符串列表，在原有slice上返回不包含空字符串的列表， 如：
func noEmpty(data []string) []string {
	out := data[:0] //在原切片上截取一个长度为0的切片 make([]string,0)
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}
	return out
}

//不使用append直接在原切片上操作
func noEmpty2(data []string) []string {
	//计数器
	i := 0
	for _, str := range data {
		if str != "" {
			//不等于空的话从第一个开始写入
			data[i] = str
			//每符合一个条件加一
			i++
		}
	}
	//返回到条件结束时的切片
	return data[:i]
}
