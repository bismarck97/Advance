package main

import "fmt"

func main08() {
	ch := make(chan string) //无缓冲channel
	//len(ch)：channel中剩余未读取数据个数     cap(ch)：通道的容量
	fmt.Println("len(ch)=", len(ch), "cap(ch)=", cap(ch))
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("i = ", i, "len(ch)=", len(ch), "cap(ch)=", cap(ch))
		}
		//通知主go程打印完毕
		ch <- "子go程打印完毕"
	}()
	str := <-ch
	fmt.Println("str = ", str)
}
