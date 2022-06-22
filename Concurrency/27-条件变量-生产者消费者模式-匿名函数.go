package main

/*
//1.创建条件变量
var cond sync.Cond

func main27() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	quit := make(chan bool)          // 创建用于结束通信的 channel

	product := make(chan int, 3) // 产品区（公共区）使用channel 模拟
	cond.L = new(sync.Mutex)     // 创建互斥锁和条件变量

	//开启5个消费者协程
	for i := 0; i < 5; i++ {
		go func(in <-chan int, index int) {
			for {
				cond.L.Lock()      // 条件变量对应互斥锁加锁（与生产者是同一个）
				for len(in) == 0 { // 产品区为空 等待生产者生产
					cond.Wait() // 挂起当前协程， 等待条件变量满足，被生产者唤醒
				}
				num := <-in // 将 channel 中的数据读走 （消费）
				fmt.Printf("消费者%d 读取了%d，公共区剩余%d个数据\n", index, num, len(in))
				cond.L.Unlock()                    //解锁条件变量用的锁
				cond.Signal()                      //唤醒阻塞的生产者者
				time.Sleep(time.Millisecond * 500) //消费完 休息一会，给其他协程执行机会
			}
		}(product, i+1)
	}
	//开启3个生产者协程
	for i := 0; i < 3; i++ {
		go func(out chan<- int, index int) {
			for {
				cond.L.Lock()       // 条件变量对应互斥锁加锁
				for len(out) == 3 { // 产品区满 等待消费者消费
					cond.Wait() // 挂起当前协程， 等待条件变量满足，被消费者唤醒
				}
				num := rand.Intn(800) //产生一个随机数
				out <- num            //写入到channel中(生产)
				fmt.Printf("生产者%d 写入了%d，公共区剩余%d个数据\n", index, num, len(out))
				cond.L.Unlock()         //解锁条件变量用的锁
				cond.Signal()           //唤醒阻塞的消费者
				time.Sleep(time.Second) // 生产完休息一会，给其他协程执行机会
			}
		}(product, i+1)
	}
	<-quit //主协程阻塞 不结束
}
*/
