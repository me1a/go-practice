package main

import "fmt"

// Vertex location
type Vertex struct {
	Lat, Long float64
}

//map 在使用之前必须用 make 而不是 new 来创建；
//值为 nil 的 map 是空的，并且不能赋值。
var m map[string]Vertex

func main() {
	fmt.Println(m == nil)
	m = make(map[string]Vertex) // panic: assignment to entry in nil map
	m["bell labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["google"] = Vertex{
		37.42202, -122.08408,
	}

	c := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m, c)
	fmt.Println(m["google"].Lat)

	// 内部同名m， 不影响外边的m
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// v 是元素的值， ok 表示该值是否存在
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
