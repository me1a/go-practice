package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // & 取该变量的内存地址
	fmt.Println(*p) // * 取该内存地址存储的值
	fmt.Println(p)  // 内存地址

	*p = 21 // 改变内存地址中的值
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}
