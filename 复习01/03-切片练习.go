package main

import "fmt"

//练习1：给定一个字符串列表，在原有slice上返回不包含空字符串的列表， 如：
//{"red", "", "black", "", "", "pink", "blue"}
//——> {"red", "black", "pink", "blue"}

func main03() {
	data := []string{"red", "", "black", "", "", "pink", "blue"}
	afterData := noEmpty(data)
	fmt.Println(afterData)
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
