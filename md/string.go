package main

import "fmt"

func test() {
	// s:="\z" // unknown escape sequence
	s := "z"
	s1 := "\172"
	s2 := "\x7A"
	s3 := "\u007A"
	// s4:="\u{7A}"//non-hex character in escape sequence: {
	fmt.Println(s, s1, s2, s3)
}

func test1() {
	s := "golang在前端"
	for i, v := range s { // Unicode字符遍历
		fmt.Printf("%d: %q [% x]\n", i, v, []byte(string(v)))
	}
	fmt.Println("----------------")
	for i := 0; i < len(s); i++ { //字节数组的方式遍历
		fmt.Println(i, s[i])
	}
}

func main() {
	test()
	test1()
}
