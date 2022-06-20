package main

import (
	"fmt"
	"time"
)

//生产者
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("生产：", i*i)
		out <- i * i
	}
	close(out)
}

//消费者
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("消费者拿到：", num)
		time.Sleep(time.Second)
	}
}
func main13() {
	//ch := make(chan int)  //同步
	ch := make(chan int, 6) //异步
	go producer(ch)         //子go程 生产者
	consumer(ch)            //主go程 消费者
}
