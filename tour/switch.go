package main

import (
	"fmt"
	"runtime"
	"time"
)

func nocond() {
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}
}

func weekday() {
	fmt.Println("when`s Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today, time.Saturday)
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("too far away")
	}

}

func main() {
	fmt.Print("go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("os x")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s.", os)

	}
	weekday()
	nocond()
}
