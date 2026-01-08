package main

import (
	"fmt"
	"net"
	"sort"

	"github.com/c-robinson/iplib"
)

func main() {

	IP := net.ParseIP("192.0.2.1")
	nextIP := iplib.NextIP(IP)
	incrIP := iplib.IncrementIPBy(nextIP, 19)

	iplist := []net.IP{incrIP, nextIP, IP}

	fmt.Println(iplist)
	sort.Sort(iplib.ByIP(iplist))
	fmt.Println(iplist)
}
