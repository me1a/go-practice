package main

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)
		for {
			x, ok := <-c
			if !ok { // 通道取完之后， 停止此go程
				return
			}
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	test()
}

func test() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)
		for v := range c {
			print(v)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
}
