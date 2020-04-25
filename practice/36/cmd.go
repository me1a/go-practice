package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var age string

func init() {
	flag.StringVar(&name, "name", "everyone", "the greating object")
	flag.StringVar(&age, "age", "22", "the greating object")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s :\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println("hello ", name)
	fmt.Println("i am ", age)
}
