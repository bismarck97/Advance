package main

import (
	"fmt"
	"runtime"
	"time"
)

func main1701() {
	ch := make(chan int)    //用来进行数据通信的channel
	quit := make(chan bool) //用来判断是否退出的channel
	//ch2 := make(chan string)

	go func() { //写数据
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true //通知主go程退出
		runtime.Goexit()
	}()
	for { //主go程    读数据
		select {
		case num := <-ch: //从channel中读取数据
			fmt.Println("读到：", num)
		case <-quit:
			//break //break跳出select语句
			return //终止进程

		}
		fmt.Println("==========================")
	}
}
func main1702() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		//select每次都是随机选择一个case执行，如果两个都满足，每次执行都可能不一样
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("啥都不干")
		}
	}
}
