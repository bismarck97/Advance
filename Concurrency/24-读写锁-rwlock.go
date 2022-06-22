package main

/*
func init() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())
}

var value int //定义全局变量，模拟共享数据

func main24() {
	//quit := make(chan bool) //用于关闭主go程的channel
	//ch := make(chan int) //用于数据传递的channel

	for i := 0; i < 5; i++ {
		go readGo(i + 1)
	}
	for i := 0; i < 5; i++ {
		go writeGo(i + 1)
	}
	//<-quit
	for {

	}
}

var rwMutex sync.RWMutex //锁只有一把，2个属性
//读
func readGo(index int) {
	for {
		rwMutex.RLock() //以读模式加锁
		num := value    //阻塞
		fmt.Printf("-----%dth 读go程，读出：%d\n", index, num)
		rwMutex.RUnlock() //以读模式解锁

	}
}

//写
func writeGo(index int) {
	for {
		//生成随机数
		num := rand.Intn(1000)
		//以写模式加锁
		//rwMutex.Lock() //阻塞
		value = num
		fmt.Printf("%dth 写go程，写入：%d\n", index, num)
		time.Sleep(time.Millisecond * 300) //放大实验现象
		//rwMutex.Unlock()
	}
}
*/
