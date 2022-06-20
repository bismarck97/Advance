package main

import (
	"fmt"
	"time"
)

//全局定义channel，用来完成数据同步操作
var channel = make(chan int)

//定义一台打印机
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch) //屏幕：stdout
		time.Sleep(time.Millisecond * 500)
	}
}

//定义两个人使用打印机
func person1() { //person1先执行
	printer("hello")
	channel <- 8
}
func person2() { //person2后执行
	<-channel
	printer("world")
}
func main() {
	go person1()
	go person2()
	select {}
}
