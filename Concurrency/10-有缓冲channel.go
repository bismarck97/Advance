package main

import (
	"fmt"
)

func main10() {
	ch := make(chan int, 3) //存满三个元素之前，不会阻塞
	fmt.Println("len = ", len(ch), "cap = ", cap(ch))

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Println("子go程传输：i = ", i, "len = ", len(ch), "cap = ", cap(ch))
		}
	}()
	//time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		n := <-ch
		fmt.Println("主go程读取：n = ", n) //io操作：耗时（访问硬件）,输出比通道同步慢，不能同步
	}
}
