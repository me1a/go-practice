package main

import "fmt"

func stuck() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // 一个函数里面的defer类似栈调用

	}
	fmt.Println("done")
}

func main() {
	defer func() {
		fmt.Println("world")
	}()
	stuck() // stuck 里的defer会在下面语句之前调用
	fmt.Println("hello")
}
