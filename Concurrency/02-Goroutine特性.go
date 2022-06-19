package main

import (
	"fmt"
	"time"
)

func main02() {

	go func() { //创建一个子go程
		for i := 0; i < 5; i++ {
			fmt.Println("I'm goroutine")
			time.Sleep(time.Second)
		}
	}()
	fmt.Println("I'm main")
	/*for i := 0; i < 5; i++ { //主go程
		fmt.Println("I'm main")
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}*/
}
