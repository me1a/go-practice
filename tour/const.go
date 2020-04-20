package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "world"
	fmt.Println(World)
	fmt.Print("\nHappy", Pi, "day\n")

	const Truth = true
	fmt.Println("go rules,", Truth)
}
