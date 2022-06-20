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
func main1502() {
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

//定时器的停止和重置
func main15() {
	myTimer := time.NewTimer(time.Second * 3) //创建定时器
	myTimer.Reset(1 * time.Second)            //重置定时时长为1
	go func() {
		<-myTimer.C
		fmt.Println("子go程，定时完毕")
	}()
	//myTimer.Stop() //设置定时器停止
	for true {

	}
}
