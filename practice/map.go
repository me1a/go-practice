package main

import (
	"fmt"
	"strings"
)

// map 字典， 对应js中object， 而非js中的Map

/*
创建map的方式
*/

func main() {
	var m1 map[int]string = map[int]string{3: "fdsf"}

	m2 := map[int]string{1: "meihuan"} // 需要在函数里面

	m3 := make(map[int]string)

	m4 := make(map[int]string, 10) // 第二个参数定义容量

	fmt.Println(m1, m2, m3, m4)

	// map 可以使用len函数，不能使用cap

	m4[9] = "9"
	m4[29] = "sd"
	m3[23] = "23"

	fmt.Println(len(m3), len(m4), m3, m4)

	// 遍历
	for k, v := range m4 {
		fmt.Println(k, v)
	}
	// 取值
	fmt.Println(m4[9])  // "9"
	fmt.Println(m4[19]) // ""
	v, has := m4[10]
	fmt.Println(v, has) // "" false

	// 删除
	delete(m4, 9)
	fmt.Println(m4)

	countStr("I am a superman do I")
}

func countStr(str string) map[string]int {
	arr := strings.Fields(str)
	res := map[string]int{}
	for i := 0; i < len(arr); i++ {
		if _, has := res[arr[i]]; has {
			res[arr[i]]++
		} else {
			res[arr[i]] = 1
		}
		fmt.Println(i, arr[i], res[arr[i]])
	}
	fmt.Println(arr)
	return res
}
