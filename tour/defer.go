package main

import "fmt"

func stuck() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	defer func() {
		fmt.Println("world")
	}()
	stuck()
	fmt.Println("hello")
}
