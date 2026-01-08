package main

import (
	"fmt"
	"net"
	"net/netip"
)

func main() {

	// ---------------------------------------------------
	// 1) IPv4 string -> netip.Addr (Multicast kontrolü)
	// ---------------------------------------------------
	ipv4, err := netip.ParseAddr("224.0.0.1")
	fmt.Printf("ipv4 type is :  %T\n", ipv4)
	if err != nil {
		panic(err)
	}

	if ipv4.IsMulticast() {
		fmt.Println("224.0.0.1 is an IPv4 Multicast address")
	}

	// ---------------------------------------------------
	// 2) IPv6 string -> netip.Addr (Link-local kontrolü)
	// ---------------------------------------------------
	ipv6, err := netip.ParseAddr("FE80:F00D::1")
	if err != nil {
		panic(err)
	}

	if ipv6.IsLinkLocalUnicast() {
		fmt.Println("FE80:F00D::1 is an IPv6 Link-Local Unicast address")
	}

	// ---------------------------------------------------
	// 3) net.IP -> netip.Addr dönüşümü
	// ---------------------------------------------------
	oldIP := net.ParseIP("192.0.2.1")

	addr, ok := netip.AddrFromSlice(oldIP)
	if !ok {
		panic("invalid IP slice")
	}

	fmt.Println("Converted from net.IP:", addr.String())

	// IPv4-mapped IPv6 mi?
	fmt.Println("Is IPv4 after unmap?:", addr.Unmap().Is4())

	// ---------------------------------------------------
	// 4) CIDR / Prefix kullanımı
	// ---------------------------------------------------
	prefix := netip.MustParsePrefix("192.0.2.0/24")

	fmt.Println("Prefix address :", prefix.Addr())
	fmt.Println("Prefix length  :", prefix.Bits())

	testIP1 := netip.MustParseAddr("192.0.2.18")
	testIP2 := netip.MustParseAddr("198.51.100.3")

	if prefix.Contains(testIP1) {
		fmt.Println(testIP1, "is inside", prefix)
	}

	if prefix.Contains(testIP2) {
		fmt.Println(testIP2, "is inside", prefix)
	} else {
		fmt.Println(testIP2, "is NOT inside", prefix)
	}
}
