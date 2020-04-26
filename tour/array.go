package main

import "fmt"

func test(arr [3]string) {
	arr[0] = "different_array"
	fmt.Println(arr) // [array in js]
}

func test2() {
	var arr [3]int = [3]int{1, 1, 1}     // 长度为3，类型为int的数组
	var arr1 [4]int                      // 声明未赋值， 默认为类型的零值 [0 0 0 0]
	arr2 := [4]int{2, 5}                 // [2 5 0 0]
	arr3 := [4]int{1, 3: 10}             // [1 0 0 10] 指定下标3的值为10
	arr4 := [...]int{1, 2, 3, 4, 99: 20} // [...]表示由后面的值推断数组的长度， 这里最长的下标是99，所以这里数组长度是100
	fmt.Println(arr, arr1, arr2, arr3, arr4)
}

func test3() {
	a := [2][2]int{{1, 1}, {1, 1}}
	fmt.Println(a)

	b := [2][3][4]int{
		{
			{0, 2, 3, 4},
			{1, 2, 4, 5},
			{2, 2, 3, 2},
		},
		{
			{0, 2, 3, 4},
			{1, 2, 4, 5},
			{2, 2, 3, 2},
		},
	} //
	fmt.Println(b)

}

func test4() {
	a := [3]int{1, 1, 1}
	b := [3]int{1, 1, 1}
	fmt.Println(a == b)
}

func main() {
	var a [5]string
	a[0] = "hello"
	a[1] = "world"
	a[2] = "world"
	a[3] = "world"
	a[4] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var eg [3]string = [3]string{"array", "in", "go"}
	test(eg)
	fmt.Println(eg)
	test2()
	test3()
	test4()
}
