package main

import (
	"fmt"
	"math"
)

// Vertex some
type Vertex struct {
	X, Y float64
}

// Abs for sqrt
// 如果结构体相当于js中的对象， 那么给该对象添加方法， 就是在func关键字和函数名指定 一个结构体
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// MyFloat d
// 你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体
// 但是，不能对来自其他包的类型或基础类型定义方法。
type MyFloat float64

// Abs d
func (f MyFloat) Abs() float64 {
	if f < MyFloat {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2())
}
