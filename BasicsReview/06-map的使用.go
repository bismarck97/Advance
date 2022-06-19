package main

import "fmt"

func main0601() {
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
func main0602() {
	m := map[int]string{1: "张三", 2: "李四", 3: "王五"}
	//遍历map  如果只需要key可以默认忽略不写value
	for k, v := range m {
		fmt.Printf("key:%d-->value:%q\n", k, v)
	}
}
func main0603() {
	//判断map中的key是否存在
	m := map[int]string{1: "张三", 2: "李四", 3: "王五"}

	if v, ok := m[1]; ok { //m[下标]返回两个值，第一个是value，第二个是bool类型代表key是否存在
		fmt.Println("value = ", v, " ok = ", ok)
	} else {
		fmt.Println("value = ", v, " ok = ", ok)
	}
}
