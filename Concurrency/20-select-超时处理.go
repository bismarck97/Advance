package main

import (
	"fmt"
	"time"
)

func main20() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() { //子go程获取数据
		for {
			select {

			case num := <-ch:
				fmt.Println("num = ", num)
			//超时处理
			case <-time.After(time.Second * 3):
				quit <- true
				goto lable
				//runtime.Goexit()
				//return
			}
		}
		//lable必须在函数内
	lable:
		fmt.Println("break to lable---------")
	}()
	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}
	<-quit //主go程，阻塞等待子go程通知，退出
	fmt.Println("finish!!!")
}
