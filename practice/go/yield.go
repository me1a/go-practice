package main

import (
	"fmt"
	"sync"
	"time"
)

func test() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Wait()
		println("wait exit go1 ")
	}()

	go func() {
		time.Sleep(time.Second)
		println("wait exit go2")
		wg.Done()
	}()

	wg.Wait()
	println("main exit")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			time.Sleep(time.Second)
			fmt.Println(id, " done")
		}(i)
	}
	fmt.Println("start")
	wg.Wait()
	test()
	fmt.Println("done")

}
