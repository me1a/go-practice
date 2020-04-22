package main

type Vertex struct {
	X, Y float64
}

// Scale pointer
// 有两个原因需要使用指针
// 1.避免在每个方法调用中拷贝值
// 2.方法可以修改接收者指向的值
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Scale1 no pointer
func (v Vertex) Scale1(f float64) {
	v.X = v.X / f
	v.Y = v.Y / f
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5) // c => 15, 20

	c := &Vertex{3, 4}
	c.Scale1(5) // c =>   3, 4
}
