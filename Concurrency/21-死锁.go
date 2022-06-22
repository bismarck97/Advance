package main

import "fmt"

//死锁1
func main2101() {
	ch := make(chan int)
	ch <- 789
	//写端阻塞，之后的代码无法执行，造成死锁,channel应该至少两个go程中使用
	num := <-ch
	fmt.Println("num: ", num)
}

//死锁2
func main2102() {
	ch := make(chan int)
	//num := <-ch //没有写入数据先进行读数据，会造成读端阻塞，之后的代码无法执行
	//fmt.Println("num = ", num)
	go func() {
		ch <- 789
	}()
	num := <-ch
	fmt.Println("num = ", num)
}

//死锁3
func main2103() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	//掌握着A条件的时候想要B条件
	go func() { //子go程
		for {
			select {
			case num := <-ch1: //监听ch1读事件
				ch2 <- num //读到数据向ch2中写入
			}
		}
	}()
	//掌握着B条件的时候想要A条件
	for { //主go程
		select {
		case num := <-ch2: //监听ch2读事件
			ch1 <- num //读到数据向ch1中写入
		}
	}
}
