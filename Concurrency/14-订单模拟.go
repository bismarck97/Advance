package main

import "fmt"

// OrderInfo 创建结构体类型OrderInfo，只有一个id 成员
type OrderInfo struct {
	id int
}

//生产者
func producer2(out chan<- OrderInfo) {
	//生产订单传送
	for i := 0; i < 10; i++ { //循环生成10份订单
		out <- OrderInfo{i + 1} //写入channel
	}
	close(out) //写完，关闭channel
}

//消费者
func consumer2(in <-chan OrderInfo) {
	for order := range in {
		fmt.Println("订单id为：", order.id)
	}
}
func main14() {
	//定义一个channel
	ch := make(chan OrderInfo)
	go producer2(ch) //子go程，传只写channel
	consumer2(ch)    //主go程，传只读channel
}
