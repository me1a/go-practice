package main

func test() {
	c := make(chan int, 2)
	println("test")
	c <- 1
	c <- 2
	c <- 2    // 调试模式执行到了这会一直挂起， 等待队头出列
	print(23) // 不会执行
}

func main() {
	// 队列， 不超过缓冲区继续执行， 直到达到缓冲大小， 如果 不消费， 继续卡住
	c := make(chan int, 3)

	c <- 1
	c <- 2

	println(<-c)
	c <- 3
	println(<-c)
	println(<-c)
	test()
}
