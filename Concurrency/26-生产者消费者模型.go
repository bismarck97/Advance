package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Producer(out chan<- int, index int) {
	for i := 0; i < 50; i++ {
		num := rand.Intn(800)
		fmt.Printf("生产者%d，生产：%d\n", index, num)
		out <- num
	}
	close(out)
}
func Consumer(in <-chan int, index int) {
	for num := range in {
		fmt.Printf("--------消费者%d，消费：%d\n", index, num)
	}
}
func main26() {
	product := make(chan int)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		go Producer(product, i+1)
	}
	for i := 0; i < 5; i++ {
		go Consumer(product, i+1)
	}
	for {

	}
}
