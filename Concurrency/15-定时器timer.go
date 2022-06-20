package main

import (
	"fmt"
	"time"
)

func main1501() {
	fmt.Println("当前时间", time.Now())
	//创建一个定时器
	myTimer := time.NewTimer(time.Second * 2)
	nowTime := <-myTimer.C //chan 类型
	fmt.Println("现下时间：", nowTime)
}

//3种定时方法
func main() {
	//1.sleep
	//time.Sleep(time.Second)
	//2.Timer.C
	//myTimer := time.NewTimer(time.Second * 2) //创建定时器，指定定时时长
	//nowTime := <-myTimer.C                    //定时满，系统自动写入系统世界
	//fmt.Println("现下时间：", nowTime)
	//3.time.After
	fmt.Println("当前时间：", time.Now())
	nowTime2 := <-time.After(time.Second * 2)
	fmt.Println("现下时间：", nowTime2)
}
