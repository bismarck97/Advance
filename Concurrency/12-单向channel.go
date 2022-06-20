package main

import "fmt"

func main1201() {
	ch := make(chan int) //双向channel

	var sendCh chan<- int = ch
	sendCh <- 789 //没有人读会造成阻塞，死锁
	//num := <-sendCh //无效参数

	var recvCh <-chan int = ch
	num := <-recvCh
	fmt.Println("num = ", num)

	//反向赋值
	//var ch2 chan int=sendCh   不能把读 写的单向chan赋值给双向chan
}
func send(out chan<- int) {
	out <- 89
	close(out)
}
func recv(in <-chan int) {
	n := <-in
	fmt.Println("读到", n)
}
func main1202() {
	ch := make(chan int) //双向channel
	go func() {
		send(ch) //双向channel转为 写channel
	}()
	recv(ch)
}
