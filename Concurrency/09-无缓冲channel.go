package main

import (
	"fmt"
	"time"
)

func main09() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子go程写，i = ", i)
			ch <- i //ch <- 0
		}
	}()
	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("主go程读，num = ", num) //io操作：耗时（访问硬件）
	}
}
