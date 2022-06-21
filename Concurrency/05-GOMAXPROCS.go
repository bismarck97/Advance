package main

import (
	"fmt"
	"runtime"
)

func main0501() {

	n := runtime.GOMAXPROCS(1) //将cpu设置为单核
	fmt.Println("n = ", n)
	n = runtime.GOMAXPROCS(2) //将cpu设置为双核
	fmt.Println("n = ", n)
	n = runtime.GOMAXPROCS(2) //将cpu设置为双核
	fmt.Println("n = ", n)
	for {
		go fmt.Print(0) //子go程
		fmt.Print(1)    //主go程
	}
}

func a() {
	for i := 0; i < 100; i++ {
		fmt.Println("a:", i)
	}
	wg.Done()
}
func b() {
	for i := 0; i < 100; i++ {
		fmt.Println("b:", i)
	}
	wg.Done()
}
func main0502() {
	runtime.GOMAXPROCS(12) //占用两个cpu核心
	wg.Add(2)
	go a()
	go b()
	wg.Wait()

}
