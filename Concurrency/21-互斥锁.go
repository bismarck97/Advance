package main

import (
	"fmt"
	"sync"
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
func main21() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
