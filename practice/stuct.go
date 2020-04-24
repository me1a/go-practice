package main

import "fmt"

// Person 结构， 对应js中的类，可为结构指定结构方法
type Person struct { // 不能赋值
	name string
	sex  bool
	age  int
}

func main() {
	var man Person = Person{"mee", true, 25}
	fmt.Println(man)
	var woman Person = Person{
		name: "lucy",
		sex:  false,
		age:  23,
	}
	fmt.Println(woman)
	fmt.Println(woman.name)

	var tmp = woman // 值传递
	fmt.Println(tmp)
	tmp.name = "harry" // 不改变原有的值
	fmt.Println(tmp, woman)
}
