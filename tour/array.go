package main

import "fmt"

func main() {
	var a [5]string
	a[0] = "hello"
	a[1] = "world"
	a[2] = "world"
	a[3] = "world"
	a[4] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}
