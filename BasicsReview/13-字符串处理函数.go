package main

import (
	"fmt"
	"strings"
)

func main13() {
	str := "I love my work and I love my family too"
	//字符串按 指定分隔符拆分
	s1 := strings.Split(str, " ")
	for _, v := range s1 {
		fmt.Println(v)
	}
	//字符串按 空格拆分
	s2 := strings.Fields(str)
	fmt.Println(s2)

	//判断字符串结束标记
	flg := strings.HasSuffix(str, "too")
	fmt.Println(flg)

	//判断字符串开始标记
	flg = strings.HasPrefix(str, "I")
	fmt.Println(flg)

}
