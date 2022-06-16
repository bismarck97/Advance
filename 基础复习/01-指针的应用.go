package main

import "fmt"

func test(a int) {
	var b int = 1000
	b += a
}
func main0101() {
	var a int = 10

	var p *int

	//在heap区分配一块内存给p指针
	p = new(int)           //默认类型的 默认值
	fmt.Printf("%q\n", *p) //打印Go语言格式的字符串
	fmt.Println(*p)

	p = &a

	a = 100
	fmt.Println("a = ", a)

	test(a)

	*p = 200 //借助a 变量的地址，操作a对应空间的值
	fmt.Println("a = ", a)
	fmt.Println("*p = ", *p)

	a = 1000
	fmt.Println("*p = ", *p)

	//野指针
	//var p *int = 0xff87ff87
	//fmt.Println(*p)
}
func swap(a, b *int) {
	*a, *b = *b, *a
}
func main0102() {
	a, b := 10, 20
	swap(&a, &b) //传地址值
	fmt.Println("main:", a, b)
}
