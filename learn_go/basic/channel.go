package main

/*
我觉得 channel 就是一个消息队列

通过 channel 实现 goroutine 之间的同步

1. 只能用 make 创建
2. ch <- 发送数据
3. v := <- ch 接受数据
4. 默认情况下 channel 接收和发送数据都是阻塞的
5. 因 4 channel 是一个很好的同步工具
6. 还可以定义一种有缓冲的 channel
	c := make(chan int, 2) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
*/

// 下面写一个单个生产者与单个消费者的例子
import (
	"fmt"
	"math/rand"
	"sync"
)

const length = 1

var buffer [length]int
var channel = make(chan bool)

func p(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < length; i++ {
		v := rand.Intn(length * 99)
		buffer[i] = v
		fmt.Printf("p produces %d \n", v)
		// 通知消费者
		channel <- true
	}
}
func c(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < length; i++ {
		<-channel
		fmt.Printf("c receives %d \n", buffer[i])

	}
}
func channelSync() {
	var wg sync.WaitGroup
	wg.Add(2)
	go p(&wg)
	go c(&wg)
	wg.Wait()
	fmt.Println("Finish!")
}
