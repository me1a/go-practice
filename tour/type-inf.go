package main

import "fmt"

func test() {
	type flags string
	// var f flags = "go"
	// var s string = "go"
	// fmt.Printf(f == s) // invalid operation: f == s (mismatched types flags and string)

	type str = string
	var ss str = "go"
	var sss string = "go"
	fmt.Println(ss == sss)
}

func main() {
	v := 52
	fmt.Printf("v is of type %T\n", v)
	test()
}
