package main

import "fmt"

func main() {
	var m1 map[int]string //声明map,没有空间，不能直接存储key--value
	//m1[101] = "张三"
	if m1 == nil {
		fmt.Println("map is nil")
	}

	fmt.Println("=============")
	m2 := map[int]string{}
	fmt.Println(len(m2))
	fmt.Println("m2 = ", m2)
	m2[4] = "hello"
	fmt.Println("m2 = ", m2)

	fmt.Println("=============")
	m3 := make(map[int]string)
	fmt.Println(len(m3))
	fmt.Println("m3 = ", m3)
	m3[40] = "word"
	fmt.Println("m3 = ", m3)

	fmt.Println("=============")
	m4 := make(map[int]string, 5) //len
	fmt.Println(len(m4))
	fmt.Println("m4 = ", m4)
	m4[400] = "Go"
	fmt.Println("m4 = ", m4)

	fmt.Println("=============")
	//不能在map中使用cap()

	//初始化map
	var m5 map[int]string = map[int]string{1: "张三", 2: "李四"}
	fmt.Println("m5 = ", m5)

	//map赋值
	m6 := make(map[int]string, 1)
	m6[100] = "张三"
	m6[101] = "李四"
	m6[102] = "hello"

	fmt.Println("m6 = ", m6)

	m6[102] = "aaa" //将原map中的key值为102的值替换

}
