package main

/*
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
func main23() {
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
*/
/*
func init() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())
}

func main23() {
	//quit := make(chan bool) //用于关闭主go程的channel
	ch := make(chan int) //用于数据传递的channel

	for i := 0; i < 5; i++ {
		go readGo(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go writeGo(ch, i+1)
	}
	//<-quit
	for {

	}
}

var rwMutex sync.RWMutex //锁只有一把，2个属性
//读
func readGo(in <-chan int, index int) {
	for {
		rwMutex.RLock() //以读模式加锁
		num := <-in     //阻塞
		fmt.Printf("-----%dth 读go程，读出：%d\n", index, num)
		rwMutex.RUnlock() //以读模式解锁

	}
}

//写
func writeGo(out chan<- int, index int) {
	for {
		//生成随机数
		num := rand.Intn(1000)
		//以写模式加锁
		//rwMutex.Lock() //阻塞
		out <- num
		fmt.Printf("%dth 写go程，写入：%d\n", index, num)
		time.Sleep(time.Millisecond * 300) //放大实验现象
		//rwMutex.Unlock()
	}
}
*/
