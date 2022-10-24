package main

import (
	"fmt"
)

//var done = make(chan bool)
//var msg string
//
//func aGoroutine() {
//	msg = "你好, 世界"
//	done <- true
//}
//
//func main() {
//	go aGoroutine()
//	<-done
//	println(msg)
//
//}

// close channel
//if bool
//then
//result = false

//var done = make(chan bool)
//var msg string
//
//func aGoroutine() {
//	msg = "你好, 世界"
//	done <- true
//	close(done)
//}
//
//func main() {
//	go aGoroutine()
//	<-done
//	println(msg)
//	println(<-done)
//}

// if done = make(chan int)
//then
//result = 0
//var done = make(chan int)
//var msg string
//
//func aGoroutine() {
//	msg = "你好, 世界"
//	done <- 1
//}
//
//func main() {
//	go aGoroutine()
//	<-done
//	close(done)
//	println(msg)
//	println(<-done)
//}

//设置缓存，无法实现同步 channel 不会等待
//var done = make(chan int, 1)
//var msg string
//
//func aGoroutine() {
//	msg = "你好, 世界"
//	done <- 1
//	close(done)
//
//}
//
//func main() {
//	go aGoroutine()
//	println(msg)
//	println(<-done)
//}

//var limit = make(chan int, 3)
//var work = []func(){
//	func() { println("1"); time.Sleep(1 * time.Second) },
//	func() { println("2"); time.Sleep(1 * time.Second) },
//	func() { println("3"); time.Sleep(1 * time.Second) },
//	func() { println("4"); time.Sleep(1 * time.Second) },
//	func() { println("5"); time.Sleep(1 * time.Second) },
//}
//
//func main() {
//	for _, w := range work {
//		go func(w func()) {
//			limit <- 1
//			w()
//			<-limit
//		}(w)
//	}
//	select {}
//}

//@li
// unkonw
//func main() {
//
//	demo1 := make(chan int)
//	for i := 0; i < 10; i++ {
//		go func() {
//			demo1 <- i
//		}()
//	}
//	println(<-demo1)
//	println(<-demo1)
//
//}

//func main() {
//	done := make(chan int) // 带 10 个缓存
//
//	// 开 N 个后台打印线程
//	for i := 0; i < 10; i++ {
//		go func() {
//			fmt.Println("你好, 世界")
//			fmt.Println(<-done)
//		}()
//	}
//
//	select {}
//
//	// 等待 N 个后台线程完成
//	//for i := 0; i < 10; i++ {
//	//fmt.Println(<-done)
//	//}
//}
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	ch := make(chan int, 64) // 成果队列

	go Producer(3, ch) // 生成 3 的倍数的序列
	go Producer(5, ch) // 生成 5 的倍数的序列
	go Consumer(ch)    // 消费生成的队列
	//case v1 := <- branches[i]: c <- v1

	// 运行一定时间后退出
	//time.Sleep(5 * time.Second)
}
