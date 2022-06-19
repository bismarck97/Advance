package main

import "fmt"

//练习3：要删除slice中间的某个元素并保存原有的元素顺序， 如：
//{5, 6, 7, 8, 9} ——> {5, 6, 8, 9}

func main05() {
	s1 := []int{5, 6, 7, 8, 9}
	s2 := remove(s1, 2)
	fmt.Println(s2)
}

func remove(data []int, index int) []int {
	//从index开始，截取+1拷贝到切片中
	copy(data[index:], data[index+1:])
	return data[:len(data)-1]
}
