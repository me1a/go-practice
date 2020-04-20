package stringutil

import "testing"
import "fmt"

func TestReverse(t *testing.T){
	cases := []struct {
		in, want string
	} {
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases {
		got := Reverse(c.in)
		fmt.Println(c.in)
		fmt.Println(c.want)
		if got != c.want{
			t.Errorf(`Reverse(%q) == %q, want %q`, c.in, got, c.want)
		}
	}
}