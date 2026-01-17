package main

import (
	"fmt"
	"net"

	"github.com/c-robinson/iplib"
)

func main() {

	IP := net.ParseIP("192.0.2.1")

	nextIP := iplib.NextIP(IP)
	incrIP := iplib.IncrementIPBy(nextIP, 19)

	// Aradaki IP sayısı
	fmt.Println(iplib.DeltaIP(IP, incrIP)) // 20

	// Karşılaştırma
	fmt.Println(iplib.CompareIPs(IP, incrIP)) // -1
}
