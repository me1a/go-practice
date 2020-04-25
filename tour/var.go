package main

import "fmt"

var c, python, java bool

func main() {
	// fmt.Println(b) //undefined: b
	var i int
	fmt.Println(i, c, python, java)

	var a int    // 0
	var b string // ""
	var c bool   // false
	var d int = 3

	fmt.Println(a, b, c, d)

	var (
		x, y int    // x = y = 0
		n, m = 2, 3 //  n = 2, m = 3
	)
	fmt.Println(x, y, n, m)

	// var x = "go" // x redeclared in this block
	fmt.Println(cc)

	func() {
		// fmt.Println(dd)undefined: dd
	}()
	// dd := 13
	const (
		v = 1 << 3
		j = 1 / 2
	)
	fmt.Println(v, j)
}

var cc = 1
