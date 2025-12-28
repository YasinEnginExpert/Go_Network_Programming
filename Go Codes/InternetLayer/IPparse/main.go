package main

import (
	"fmt"
	"net"
)

func main() {
	// ParseIP returns the IP address parsed from s.
	// The string s can be in dotted decimal ("192.0.2.1") or IPv6 format ("FC02:F00D::1").
	// To4 returns the 4-byte representation of the IPv4 address ip.
	// To16 returns the 16-byte representation of the IP address ip.
	// If ip is not an IP address, To4 and To16 return nil.
	// If ip is an IPv4 address, To16 returns the 16-byte representation of the IPv4 address.
	// If ip is an IPv6 address, To4 returns nil.
	// The returned IP address is in 16-byte form.
	// The IPv4 address is converted to IPv6 form by prepending
	// the 12-byte prefix ::ffff:0:0/96.
	// See https://tools.ietf.org/html/rfc4291#section-2.5.5.2
	// for more information.
	// Example:

	ipv4 := net.ParseIP("192.0.2.1").To4()
	fmt.Printf("Type: %T\n", ipv4)

	ipv6 := net.ParseIP("FC02:F00D::1").To16()
	fmt.Printf("Type: %T\n", ipv6)

	// Print the results
	fmt.Println("IPv4:", ipv4)
	fmt.Println("IPv6:", ipv6)
	fmt.Println("Length of IPv4 byte slice:", len(ipv4))
	fmt.Println("Length of IPv6 byte slice:", len(ipv6))

	// Print each byte in binary, decimal, and hexadecimal format
	// for both IPv4 and IPv6 addresses
	fmt.Println("Bytes of IPv4 address:")
	for i, b := range ipv4 {
		fmt.Printf("Byte %d: %08b (decimal=%d, hex=%02x)\n",
			i, b, b, b)
	}

	fmt.Println("Bytes of IPv6 address:")
	for i, b := range ipv6 {
		fmt.Printf("Byte %d: %08b (decimal=%d, hex=%02x)\n",
			i, b, b, b)
	}

}
