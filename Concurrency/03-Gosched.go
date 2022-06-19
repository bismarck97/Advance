package main

import (
	"fmt"
	"runtime"
)

func main03() {
	go func() {
		for {
			fmt.Println("this is goroutine test")
			//time.Sleep(100 * time.Microsecond)
		}
	}()
	for {
		runtime.Gosched() //出让当前cpu时间片
		fmt.Println("this is main test")
		//time.Sleep(100 * time.Microsecond)
	}
}
