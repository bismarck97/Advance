package main

import (
	"fmt"
	"unsafe"
)

func test03(p *person) {

	p.name = "李四"
	p.age = 20
	p.sex = "女"
	fmt.Println("test size:", unsafe.Sizeof(p))
}
func main11() {
	var p1 *person = &person{name: "张三", age: 18}
	fmt.Println(p1)
	fmt.Println("==========================")
	var p2 *person = newPerson("李四", "男", 20)
	fmt.Println(p2)
	fmt.Println("==========================")
	p3 := newPerson("王五", "男", 21)
	fmt.Println(p3)
	fmt.Println("==========================")
	var p4 *person
	p4 = new(person)
	p4 = &person{name: "张三", age: 18}
	fmt.Println(p4)
	fmt.Println("==========================")
	fmt.Printf("p4 = %p\n", p4)
	fmt.Printf("&p4.name = %p\n", &p4.name)
	test03(p4)
	fmt.Println(p4)
	fmt.Println("main size:", unsafe.Sizeof(p4))
}
