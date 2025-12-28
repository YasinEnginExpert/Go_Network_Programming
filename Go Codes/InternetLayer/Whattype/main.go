package main

import (
	"fmt"
	"net/netip"
)

// analyzeIPv6 analyzes and prints the properties of the given IPv6 address.
// It checks for various types such as Global Unicast, Link-Local, Loopback, Unspecified,
// Unique Local (ULA), Multicast, and IPv4-Embedded (4in6).
// analyzeIPv4 analyzes and prints the properties of the given IPv4 address.
// It checks for properties such as Private, Loopback, Unspecified, and Multicast.
// The results are printed to the console.
func analyzeIPv6(addrStr string) {
	ip := netip.MustParseAddr(addrStr)

	fmt.Println("IPv6 Adres :", ip)

	fmt.Println("  Global Unicast        :", ip.IsGlobalUnicast())
	fmt.Println("  Link-Local Unicast    :", ip.IsLinkLocalUnicast())
	fmt.Println("  Loopback              :", ip.IsLoopback())
	fmt.Println("  Unspecified           :", ip.IsUnspecified())
	fmt.Println("  Unique Local (ULA)    :", ip.IsPrivate())
	fmt.Println("  Multicast             :", ip.IsMulticast())
	fmt.Println("  IPv4 Embedded (4in6)  :", ip.Is4In6())

	fmt.Println("--------------------------------------")
}

// analyzeIPv4 analyzes and prints the properties of the given IPv4 address.
// It checks for properties such as Private, Loopback, Unspecified, and Multicast.
// The results are printed to the console.
// analyzeIPv4 analyzes and prints the properties of the given IPv4 address.
// It checks for properties such as Private, Loopback, Unspecified, and Multicast.

func analyzeIPv4(addrStr string) {
	ip := netip.MustParseAddr(addrStr)
	fmt.Println("IPv4 Address :", ip)
	fmt.Println("  Private        :", ip.IsPrivate())
	fmt.Println("  Loopback       :", ip.IsLoopback())
	fmt.Println("  Unspecified    :", ip.IsUnspecified())
	fmt.Println("  Multicast      :", ip.IsMulticast())
	fmt.Println("--------------------------------------")
}

func main() {
	// Test with various IPv6 addresses
	ipv6Addresses := []string{
		"2001:4860:4860::8888", // Global Unicast
		"fe80::1",              // Link-Local
		"::1",                  // Loopback
		"::",                   // Unspecified
		"fd12:3456:789a::1",    // Unique Local (Private)
		"ff02::1",              // Multicast (All Nodes)
		"ff02::2",              // Multicast (All Routers)
		"ff02::1:ffab:cdef",    // Solicited-Node Multicast
		"::ffff:192.0.2.1",     // IPv4-Embedded IPv6
	}

	// Analyze each IPv6 address
	for _, ip := range ipv6Addresses {
		analyzeIPv6(ip)
	}

	// Test with various IPv4 addresses
	ipv4Addresses := []string{
		"192.168.1.10",  // IPv4 Private
		"8.8.8.8",       // IPv4 Public
		"10.0.0.5",      // IPv4 Private
		"172.16.0.1",    // IPv4 Private
		"127.0.0.1",     // IPv4 Loopback
		"0.0.0.0",       // IPv4 Unspecified
		"213.123.45.67", // IPv4 Public
		"4.4.4.4",       // IPv4 Public

	}

	// Analyze each IPv4 address
	for _, ip := range ipv4Addresses {
		analyzeIPv4(ip)
	}
}
