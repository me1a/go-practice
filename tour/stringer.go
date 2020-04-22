package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip *IPAddr) String() string {
	s := make([]string, len(ip))
	for i := range ip {
		s[i] = strconv.Itoa(int(ip[i]))
	}
	return strings.Join(s, ".")
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Println(n, a.String())
		fmt.Printf("%v: %v\n", n, a)
	}
}
