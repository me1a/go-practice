package main

import (
	"runtime"
	"time"
)

func main() {

	n := runtime.GOMAXPROCS(1)
	println(n)
	println(runtime.GOROOT())
	println(runtime.NumCPU())

	for {
		go func() {
			print(1)
			// runtime.Gosched()
		}()
		print(0)
		time.Sleep(time.Second)
	}

}
