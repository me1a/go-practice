package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32}
var obj = map[string]string{
	"name": "meihuan",
	"age":  "28",
	"sex":  "man",
}

// range 格式可以对 slice 或者 map 进行迭代
func main() {

	for i, v := range pow {
		fmt.Println(i, v)
	}
	for i, v := range obj {
		fmt.Println(i, v)
	}
}
