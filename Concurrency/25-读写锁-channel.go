package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())
}

func main25() {
	//quit := make(chan bool) //用于关闭主go程的channel
	ch := make(chan int) //用于数据传递的channel

	for i := 0; i < 5; i++ {
		go readGo(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go writeGo(ch, i+1)
	}
	//<-quit
	for {

	}
}

//var value int //定义全局变量，模拟共享数据
//写
func writeGo(out chan<- int, index int) {
	for {
		//生成随机数
		num := rand.Intn(1000)
		out <- num //写入channel
		fmt.Printf("%dth 写go程，写入：%d\n", index, num)
		time.Sleep(time.Millisecond * 300) //放大实验现象

	}
}

//读
func readGo(in <-chan int, index int) {
	for {
		num := <-in //从channel中读取数据
		fmt.Printf("-----%dth 读go程，读出：%d\n", index, num)
	}
}
