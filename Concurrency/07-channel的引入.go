package main

import "fmt"

/**
//全局定义channel，用来完成数据同步操作
var channel = make(chan int)

//定义一台打印机
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch) //屏幕：stdout
		time.Sleep(time.Millisecond * 500)
	}
}

//定义两个人使用打印机
func person1() { //person1先执行
	printer("hello")
	channel <- 8
}
func person2() { //person2后执行
	<-channel
	printer("world")
}
func main07() {
	go person1()
	go person2()
	select {}
}
**/
func main07() {
	//通道有发送(send)、接收(receive)和关闭(close)三种操作
	//发送和接收都使用<-符号
	//var ch1 chan int //引用类型，需要初始化之后才能使用
	//ch1 = make(chan int, 1)
	//ch1 := make(chan int)//无缓冲区通道，又称为同步通道
	ch1 := make(chan int, 1) //带缓冲区通道，又称为异步通道
	ch1 <- 20                //发送值，无缓冲区不能暂存值
	num := <-ch1
	fmt.Println(num)
	//len(ch1) //取通道中元素的数量
	//cap(ch1) //拿到通道的容量
	close(ch1)

}

/*两个goroutine，两个channel
1.生成0-100的数字发送到ch1
2.从ch1取出数据计算它的平方，把结果发送到ch2中
*/

func main0702() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	//单向通道 ch单向写
	go func(ch chan<- int) {
		for i := 0; i < 100; i++ {
			//1.生成0-100的数字发送到ch1
			ch <- i
		}
		//<-ch
		close(ch1) //使用完通道应该关闭掉，数据会存储到缓冲区中
	}(ch1)
	//ch1单向读，ch2单向写
	go func(ch1 <-chan int, ch2 chan<- int) {
		//从ch1取出数据计算它的平方，把结果发送到ch2中
		//取通道值的两种方式，推荐第二种
		//第一种
		//for {
		//	if tmp, ok := <-ch1; ok {
		//		ch2 <- tmp * tmp
		//	} else {
		//		break
		//	}
		//}
		//第二种
		for tmp := range ch1 {
			ch2 <- tmp * tmp
		}

		close(ch2)
	}(ch1, ch2)
	//主go程使用遍历取出ch2的值
	for num := range ch2 {
		fmt.Println(num)
	}
}
