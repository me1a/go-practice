package main

import "fmt"

func test1() {
	p := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println("p[1:4] ==", p[1:4])

	fmt.Println("p[:3]==", p[:3])
	fmt.Println("p[4:]==", p[4:])
}

func test2() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
	e := []int{}
	printSlice("e", e)
	var f []int
	printSlice("f", f)
	println(e == nil)
	println(f == nil)

}
func test3() {
	println("-------------------")
	var a []int
	printSlice("a", a)
	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	p := []int{2, 3, 4, 5, 6, 7}
	fmt.Println("p==", p)

	for i := 0; i < len(p); i++ {
		// fmt.Printf("p[%d] == %d\n", i, p[i])
	}
	test1()
	test2()
	test3()
}
