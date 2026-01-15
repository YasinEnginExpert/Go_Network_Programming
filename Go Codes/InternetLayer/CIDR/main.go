// Package main demonstrates CIDR (Classless Inter-Domain Routing) operations in Go.
//
// CIDR is a method for allocating IP addresses and routing IP packets. It replaced
// the original class-based addressing system (Class A, B, C) with a more flexible
// approach using variable-length subnet masking.
//
// Key Concepts:
//
//   - CIDR Notation: IP address followed by a slash and prefix length (e.g., 192.168.1.0/24)
//   - Prefix Length: Number of leading bits that form the network portion
//   - Subnet Mask: Derived from prefix length, indicates network vs host bits
//
// This program demonstrates:
//  1. Creating CIDR masks for IPv4 and IPv6
//  2. Parsing CIDR notation strings
//  3. Extracting network and host information
//
// Reference: RFC 4632 - Classless Inter-domain Routing (CIDR)
//
// Usage:
//
//	go run main.go
package main

import (
	"fmt"
	"net"
)

func main() {
	// ============================================================
	// PART 1: Creating CIDR Masks
	// ============================================================
	//
	// net.CIDRMask creates a subnet mask from prefix length and total bits.
	// Parameters: (ones, bits) where:
	//   - ones: number of 1 bits (network portion)
	//   - bits: total bits in the mask (32 for IPv4, 128 for IPv6)

	fmt.Println("=== CIDR Mask Creation ===")

	// IPv4 /31 mask (point-to-point link, only 2 usable addresses)
	// Binary: 11111111.11111111.11111111.11111110
	mask31 := net.CIDRMask(31, 32)
	fmt.Printf("IPv4 /31 mask: %d.%d.%d.%d\n", mask31[0], mask31[1], mask31[2], mask31[3])
	fmt.Printf("               Binary: %08b.%08b.%08b.%08b\n", mask31[0], mask31[1], mask31[2], mask31[3])

	// IPv4 /24 mask (standard Class C size network)
	// Binary: 11111111.11111111.11111111.00000000
	// This provides 256 addresses (254 usable for hosts)
	mask24 := net.CIDRMask(24, 32)
	fmt.Printf("IPv4 /24 mask: %d.%d.%d.%d\n", mask24[0], mask24[1], mask24[2], mask24[3])
	fmt.Printf("               Binary: %08b.%08b.%08b.%08b\n", mask24[0], mask24[1], mask24[2], mask24[3])

	// IPv6 /64 mask (standard subnet size for IPv6)
	// First 64 bits are network, remaining 64 bits are interface identifier
	mask64v6 := net.CIDRMask(64, 128)
	fmt.Printf("IPv6 /64 mask: %x\n", mask64v6)

	fmt.Println()

	// ============================================================
	// PART 2: Parsing CIDR Notation
	// ============================================================
	//
	// net.ParseCIDR parses a CIDR notation string and returns:
	//   - IP: The specific IP address in the CIDR notation
	//   - IPNet: The network (IP masked with subnet mask)
	//   - error: Any parsing error

	fmt.Println("=== CIDR Parsing ===")

	// Parse an IPv4 CIDR address
	// "192.0.2.1/24" means:
	//   - Host IP: 192.0.2.1
	//   - Network: 192.0.2.0/24 (256 addresses)
	ip4, ipNet4, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		fmt.Println("Error parsing IPv4 CIDR:", err)
	} else {
		fmt.Println("IPv4 CIDR: 192.0.2.1/24")
		fmt.Println("  Host IP:      ", ip4)
		fmt.Println("  Network:      ", ipNet4)
		fmt.Println("  Network IP:   ", ipNet4.IP)
		fmt.Println("  Subnet Mask:  ", ipNet4.Mask)
		ones, bits := ipNet4.Mask.Size()
		fmt.Printf("  Prefix Length: /%d (out of %d bits)\n", ones, bits)
	}

	fmt.Println()

	// Parse an IPv6 CIDR address
	// "2001:db8::1/64" is a documentation address (RFC 3849)
	ip6, ipNet6, err := net.ParseCIDR("2001:db8::1/64")
	if err != nil {
		fmt.Println("Error parsing IPv6 CIDR:", err)
	} else {
		fmt.Println("IPv6 CIDR: 2001:db8::1/64")
		fmt.Println("  Host IP:    ", ip6)
		fmt.Println("  Network:    ", ipNet6)
		fmt.Println("  Network IP: ", ipNet6.IP)
		ones, bits := ipNet6.Mask.Size()
		fmt.Printf("  Prefix Length: /%d (out of %d bits)\n", ones, bits)
	}

	fmt.Println()

	// ============================================================
	// PART 3: Checking IP Containment
	// ============================================================
	//
	// IPNet.Contains() checks if an IP address belongs to a network

	fmt.Println("=== IP Containment Check ===")

	testIPs := []string{"192.0.2.50", "192.0.3.1", "10.0.0.1"}
	for _, testIP := range testIPs {
		ip := net.ParseIP(testIP)
		if ipNet4.Contains(ip) {
			fmt.Printf("  %s is within %s\n", testIP, ipNet4)
		} else {
			fmt.Printf("  %s is NOT within %s\n", testIP, ipNet4)
		}
	}
}
