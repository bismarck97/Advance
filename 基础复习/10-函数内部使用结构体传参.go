package main

import (
	"fmt"
	"unsafe"
)

func test01(p person) {
	p.name = "张三"
	p.age = 18
	p.sex = "男"
	fmt.Println("test01 p1 size:", unsafe.Sizeof(p))
	fmt.Printf("%p\n", &p)
	fmt.Println("test01:p1", p)
	fmt.Println("====================")

}
func test02(p *person) {
	p.name = "李四"
	p.age = 18
	p.sex = "男"
	fmt.Println("test01 p2 size:", unsafe.Sizeof(p))
	fmt.Printf("%p\n", p)
	fmt.Println("test02:p", p)
	fmt.Println("====================")
}
func main10() {
	var tmp1 person
	test01(tmp1) //值传递，将实参的值拷贝一份给形参
	fmt.Println("main p1 size:", unsafe.Sizeof(tmp1))
	fmt.Printf("%p\n", &tmp1)
	fmt.Println("main p1", tmp1)
	fmt.Println("====================")

	tmp2 := newperson()
	test02(tmp2) //引用传递，传的是地址 形参影响实参的值
	fmt.Println("main p2 size:", unsafe.Sizeof(tmp2))
	fmt.Printf("%p\n", tmp2)
	fmt.Println("main p2", tmp2)
	fmt.Println("====================")

}
