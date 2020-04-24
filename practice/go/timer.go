package main

import (
	"fmt"
	"time"
)

// type Timer struct {
// 	C <-chan Time
// 	r runtimeTimer
// }

func timeout() {

	fmt.Println(time.Now())
	myTimer := time.NewTimer(2 * time.Second)
	nowTime := <-myTimer.C
	fmt.Println(nowTime)

	nowTime = <-time.After(2 * time.Second)
	fmt.Println(nowTime)

	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())

	myTimer = time.NewTimer(2 * time.Second)
	myTimer.Stop()
}

func interval() {
	quit := make(chan bool)
	fmt.Println("now:", time.Now())
	myTicker := time.NewTicker(time.Second)

	i := 0
	go func() {
		for {
			nowTime := <-myTicker.C // 系统内部会促发这个channel，一秒出发一次
			i++
			fmt.Println("nowTime", nowTime)
			if i == 8 {
				quit <- true
				break
				// return
				// runtime.Goexit()
			}
		}
	}()

	<-quit

}

func main() {
	// timeout()
	interval()
}
