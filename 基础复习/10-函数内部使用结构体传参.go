package main

import "fmt"

func test01(p person) {
	p.name = "张三"
	p.age = 18
	p.sex = "男"
	fmt.Println("test:p", p)

}
func main() {
	var tmp person
	test01(tmp)
}
