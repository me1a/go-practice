package main

import (
	"fmt"
	"time"
)

var channel = make(chan int)

func printer(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
		time.Sleep(300 * time.Millisecond)
	}
}

func person1() {
	printer("hello ")
	channel <- 8
}
func person2() {
	<-channel
	printer("world")
}

func main() {
	go person1()
	go person2()
	for {
	}
}
