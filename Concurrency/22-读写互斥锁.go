package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	n      int64
	wgs    sync.WaitGroup
	locks  sync.Mutex //互斥锁
	rwLock sync.RWMutex
)

//读写互斥锁
func read() {
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	wgs.Done()
}
func write() {
	rwLock.RLock()
	n++
	time.Sleep(time.Millisecond * 10)
	rwLock.RUnlock()
	wgs.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wgs.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wgs.Add(1)
		go write()
	}
	wgs.Wait()
	//计算耗时
	fmt.Println(time.Now().Sub(start))
}
