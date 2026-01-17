package main

import (
	"fmt"
	"net"

	"github.com/c-robinson/iplib"
)

func main() {

	n4 := iplib.NewNet4(net.ParseIP("198.51.100.0"), 24)

	fmt.Println("Total IP addresses:", n4.Count())
	fmt.Println("First three IPs:", n4.Enumerate(3, 0))
	fmt.Println("First IP:", n4.FirstAddress())
	fmt.Println("Last IP:", n4.LastAddress())
}
