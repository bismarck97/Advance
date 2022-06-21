package main

import (
	"fmt"
	"sync"
	"time"
)

func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("----正在唱歌----")
		time.Sleep(time.Microsecond * 100)
	}
}
func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("----正在跳舞----")
		time.Sleep(time.Microsecond * 100)
	}
}

//wg是一个全局变量，用来计数
var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello", i)
	wg.Done() //通知wg把计数牌-1
}

func main0101() { //开启一个主goroutine去执行goroutine去执行main函数
	//开启独立的goroutine去执行子go程函数

	go sing()
	go dance()

}
func main0102() {
	//执行顺序不固定
	wg.Add(1000) //计数牌+1
	for i := 0; i < 1000; i++ {
		go hello(i)
	}
	fmt.Println("hello main")
	//time.Sleep(time.Second)   不建议
	wg.Wait() //阻塞，等待wg计数牌为0
}
func main0103() {
	wg.Add(1000) //计数牌+1000

	//for i := 0; i < 1000; i++ {
	for i := 0; i < 1000; i++ {
		//匿名函数开启goroutine
		go func(i int) {
			//如果i不传入的话，匿名函数就是一个闭包，变量内部i会从外面找，可能外部i已经走了很多步，匿名函数才刚刚获取到i，会造成i重复现象
			fmt.Println("hello ", i)
			wg.Done() //通知wg把计数牌-1
		}(i)
	}
	fmt.Println("hello main")
	wg.Wait()
}
