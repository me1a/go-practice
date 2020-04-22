package main

import (
	"fmt"
	"os"
)

func create() {
	f, err := os.Create("./createByGo.x")
	if err == nil {
		fmt.Println("create successful")
	}
	if i, err := f.WriteString("hello world"); err == nil {
		fmt.Println("write successful, size", i)
	}

	defer f.Close()
}

func main() {
	create()
}
