package main

import (
	"fmt"
	"runtime"
	"time"
)

func main17() {
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
