package main

import (
	"fmt"
	"time"
)

func main16() {
	quit := make(chan bool) //创建一个判断是否终止的channel
	fmt.Println("now：", time.Now())
	myTicker := time.NewTicker(time.Second) //定义一个周期定时器
	i := 0
	go func() {
		for {
			nowTime := <-myTicker.C
			i++
			fmt.Println("nowTime", nowTime)
			if i == 8 {
				quit <- true //解除主go程阻塞
				break        //return //runtime.Goexit
			}
		}

	}()
	<-quit //子go程  循环获取 <-myTicker.C期间，一直阻塞
}
