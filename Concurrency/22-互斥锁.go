package main

import (
	"fmt"
	"sync"
	"time"
)

//多个goroutine并发操作全局变量x
var (
	x int64
	//通过加锁使goroutine实现有序的操作
	lock sync.Mutex //互斥锁
)

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock() //加锁
		x++
		lock.Unlock() //释放锁
	}
	wg.Done()
}
func main2201() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

//使用channel完成同步
/*var ch = make(chan int)

func printer(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
}

//写先读后
func person1() { //person1先执行
	printer("hello")
	ch <- 1 //写
}
func person2() { //person2后执行
	<-ch //读
	printer("world")

}

func main2202() {
	go person1()
	go person2()
	select {}
}*/

//使用 锁 完成同步 -->互斥锁
var mutex sync.Mutex //创建一个互斥量，新建的互斥锁状态为0，未加锁，锁只有一把

func printer(str string) {
	mutex.Lock() //访问共享资源之前，加锁
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock() //共享数据访问结束，解锁
}

func person1() { //person1先执行
	printer("hello")
}
func person2() { //person2后执行
	printer("world")
}
func main2203() {
	go person1()
	go person2()
	for {

	}
}
