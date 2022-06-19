package main

import (
	"fmt"
	"runtime"
)

func main04() {
	go func() {
		defer fmt.Println("a")
		test()
		fmt.Println("b")
	}()
	for { //不让主go程结束
	}
}
func test() {
	fmt.Println("c")
	//return
	runtime.Goexit() //退出当前go程
	defer fmt.Println("d")
}
