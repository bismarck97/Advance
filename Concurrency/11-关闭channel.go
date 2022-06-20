package main

import (
	"fmt"
	"time"
)

func main11() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) //写端，写完数据主动关闭channel
		fmt.Println("子go结束")
	}()
	time.Sleep(time.Second * 5)
	//for {
	//	if num, ok := <-ch; ok {
	//		fmt.Println("读到数据：", num)
	//	} else {
	//		n := <-ch
	//		fmt.Println("关闭后：", n)
	//		break
	//	}
	//}
	for i := range ch {
		fmt.Println("读到数据：", i)
	}
}
