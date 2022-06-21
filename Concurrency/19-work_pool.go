package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, res chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, job)
		res <- job * 2
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("worker:%d stop job:%d\n", id, job)
	}
}
func main19() {
	jobs := make(chan int, 100)
	res := make(chan int, 100)
	//开启三个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, res)
	}
	//发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	//输出结果
	for i := 0; i < 5; i++ {
		ret := <-res
		fmt.Println(ret)
	}
}
