package main

import (
	"fmt"
	"net"
)

func main() {
	// CIDRMask returns an IP mask consisting of 'ones' 1 bits followed by
	// 'zeros' 0 bits. The total number of bits is specified by bits.
	// For example, CIDRMask(24, 32) returns the IPv4 mask
	// 11111111.11111111.11111111.00000000
	// which is equivalent to

	// For IPv4
	mask := net.CIDRMask(31, 32)
	fmt.Printf("IPv4 /31 mask: %08b\n", mask)

	// For IPv6
	mask6 := net.CIDRMask(64, 128)
	fmt.Printf("IPv6 /64 mask: %08b\n", mask6)

	// ParseCIDR parses a CIDR notation IP address and prefix length,
	// like "192.0.2.1/24". It returns the IP address and the network
	// implied by the IP and prefix length.
	// For example, ParseCIDR("192.0.2.1/24") returns the IP address
	// and the network implied by the IP and prefix length.

	// For IPv4
	ip, ipNet, _ := net.ParseCIDR("192.0.2.1/24")
	fmt.Println("IPv4 Network:", ipNet)
	fmt.Println("IPv4 IP     :", ip)

	// For IPv6
	ip6, ipNet6, _ := net.ParseCIDR("2001:db8::1/64")
	fmt.Println("IPv6 Network:", ipNet6)
	fmt.Println("IPv6 IP     :", ip6)
}
