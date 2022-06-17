package main

import (
	"fmt"
	"strings"
)

//练习3：	封装 wcFunc() 函数。接收一段英文字符串str。返回一个map，记录str中每个“词”出现次数的。
//如："I love my work and I love my family too"
//输出：
//family : 1
//too : 1
//I : 2
//love : 2
//my : 2
//work : 1
//and : 1
//提示：使用 strings.Fields() 函数可提高效率。

func main() {
	str := "I love my work and I love my family too"
	s1 := wcFunc(str)
	for k, v := range s1 {
		fmt.Printf("%q : %d\n", k, v)
	}
}

func wcFunc(str string) map[string]int {
	s := strings.Fields(str)
	m := make(map[string]int)
	for _, value := range s {
		if _, ok := m[value]; ok {
			m[value] += 1
		} else {
			m[value] = 1
		}
	}
	return m
}
