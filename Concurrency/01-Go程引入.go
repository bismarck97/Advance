package main

import (
	"fmt"
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
func main01() {
	go sing()
	go dance()
	for {

	}
}
