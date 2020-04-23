package main

func main() {
	done := make(chan int)
	c := make(chan string)

	go func() {
		s := <-c // 再收
		println(s)
		close(done)
	}()

	c <- "hi" // 1 先发
	<-done
}
