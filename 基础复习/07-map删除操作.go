package main

import "fmt"

//map做函数参数，返回值，传引用
func main07() {
	m1 := map[int]string{1: "张三", 2: "李四", 3: "王五"}
	fmt.Println(m1)
	mapDelete(m1, 3)
	fmt.Println(m1)
}

func mapDelete(m map[int]string, key int) {
	delete(m, key) //删除m中键值为key的map元素
}
