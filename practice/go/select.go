package main

import (
	"fmt"
	"runtime"
	"time"
)

// select
// 类似switch， 针对channel
// 命中多个， 选择一个来执行
// 如果没有命中，
//     如果有default， 那么会执行default， 同时程序执行会从select语句中恢复
//     如果没有有default，那么select将会被阻塞，直到至少有一个通信进行下去
// 自身不带循环机制，需要借助外层循环

func main() {

	quit := make(chan bool)

	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch) // 关闭之后还能读取
		quit <- true
		runtime.Goexit()
	}()

	ch1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			time.Sleep(time.Second)
		}
		close(ch1)
		quit <- true
		runtime.Goexit()
	}()

	for {
		select {
		case num := <-ch:
			fmt.Println("read ch:", num)
		case num := <-ch1:
			fmt.Println("read ch1:", num)
		case <-quit:
			fmt.Println("============")
			return
		}
	}

}
