package main

import "fmt"

type person struct {
	name string
	sex  string
	age  int
}

//构造函数
func newPerson(name, sex string, age int) *person {
	return &person{
		name: name,
		age:  age,
		sex:  sex,
	}
}
func main09() {
	//利用构造函数创建结构体指针变量
	p := newPerson("张三", "男", 18)
	fmt.Printf("%T\n", p)
	fmt.Println(*p)

	//普通方法创建函数
	var man person
	man.name = "李四"
	man.sex = "男"
	man.age = 99
	fmt.Printf("%T\n", man)
	fmt.Println(man)
	fmt.Println("======================")
	//普通结构体比较的是里面的值
	var p1 person = person{"张三", "男", 18}
	var p2 person = person{"张三", "男", 18}
	var p3 person = person{"李四", "男", 118}
	fmt.Println("p1 == p2 ?", p1 == p2)
	fmt.Println("p1 == p3 ?", p1 == p3)
	//打印结构体内存地址
	fmt.Printf("p1: %p\n", &p1)
	fmt.Printf("p2: %p\n", &p2)
	fmt.Printf("p3: %p\n", &p3)
	fmt.Println("======================")
	//结构体指针比较的是内存地址
	p5 := newPerson("张三", "男", 18)
	p6 := newPerson("张三", "男", 18)
	fmt.Printf("p5: %p\n", &p5)
	fmt.Printf("p6: %p\n", &p6)
	fmt.Println("p5 == p6 ?", p5 == p6)
	fmt.Println("======================")
	//相同类型结构体赋值
	var tmp person
	fmt.Println("tmp:", tmp)
	tmp = p1
}
