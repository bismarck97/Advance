package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond

func main() {
	//创建随机数种子
	rand.Seed(time.Now().UnixNano())
	//创建用于阻塞主go程的chan
	quit := make(chan bool)

	ch := make(chan int, 3)
	cond.L = new(sync.Mutex) //指定条件变量用的锁

	//创建5个消费者
	for i := 0; i < 5; i++ {
		go Consumer2(ch, i+i)
	}
	//创建三个生产者
	for i := 0; i < 3; i++ {
		go Producer2(ch, i+i)
	}
	<-quit //主go程阻塞 不结束

}
func Producer2(out chan<- int, index int) {
	for {
		cond.L.Lock()       //给公共区加锁
		for len(out) == 3 { //判断是否达到阻塞条件
			cond.Wait()
		}
		//访问公共区
		num := rand.Intn(800) //生成随机数
		out <- num            //将随机数写入公共区
		fmt.Printf("生产者%d 写入了%d，公共区剩余%d个数据\n", index, num, len(out))
		//解锁
		cond.L.Unlock()
		cond.Signal()           //唤醒消费者
		time.Sleep(time.Second) // 生产完休息一会，给其他协程执行机会
	}
}
func Consumer2(in <-chan int, index int) {
	for {
		cond.L.Lock() //给公共区加锁
		for len(in) == 0 {
			cond.Wait()
		}
		//访问公共区
		num := <-in
		fmt.Printf("消费者%d 读取了%d，公共区剩余%d个数据\n", index, num, len(in))
		cond.L.Unlock()
		cond.Signal()                      //唤醒生产者
		time.Sleep(time.Millisecond * 500) // 生产完休息一会，给其他协程执行机会,读一般比写快
	}
}
