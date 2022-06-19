package main

import (
	"fmt"
	"runtime"
)

func main05() {

	n := runtime.GOMAXPROCS(1) //将cpu设置为单核
	fmt.Println("n = ", n)
	n = runtime.GOMAXPROCS(2) //将cpu设置为双核
	fmt.Println("n = ", n)
	n = runtime.GOMAXPROCS(2) //将cpu设置为双核
	fmt.Println("n = ", n)
	for {
		go fmt.Print(0) //子go程
		fmt.Print(1)    //主go程
	}
}
