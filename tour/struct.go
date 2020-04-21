package main

import "fmt"

// Vertex is a struct
// 结构体是值传递
// 作为函数参数传递时， 会复制一份
// 如果想在一个包中访问另一个包中结构体的字段，则必须是大写字母开头的变量，即可导出的变量
type Vertex struct {
	X int
	Y int
}

// more info https://juejin.im/post/5ca2f37ce51d4502a27f0539
func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 3
	fmt.Println(v.X)

	p := &v //
	fmt.Println(*p, p)
	p.Y = 1e9
	fmt.Println(v)

	c := Vertex{X: 1, Y: 2}
	fmt.Println(c)
}
