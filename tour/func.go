package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

// func(int) int 是一个整体， 表示adder的返回值是函数，该函数参数是int，返回值是int
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	fresh := 1
	fresher := 1

	count := 0
	return func() int {
		count++
		if count <= 2 {
			return 1
		}
		res := fresh + fresher
		fresh, fresher = fresher, res
		return res

	}
}

func main() {
	fmt.Println(add(32, 23))
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*1))
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
