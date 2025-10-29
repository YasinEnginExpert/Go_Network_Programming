package main

import (
	"fmt"
	"net"
)

func main() {
	// Define an IP and subnet mask

	ip := net.ParseIP("192.168.10.1")      // String nesnesini gerçek IP adresine dönüştürür
	mask := net.IPv4Mask(255, 255, 255, 0) // Subnet Maskı oluşturur

	// Apply the mask to find the network address

	networkaddress := ip.Mask(mask) // Bu işlem Ip adresi ve maskı and işlemie sokar ve network elde edilir
	fmt.Println("Ip address :", ip)
	fmt.Println("Subnet mask :", mask)
	fmt.Println("Network address :", networkaddress)

	// Convert to CIDR notation

	cird := fmt.Sprintf("%s/%d", networkaddress.String(), maskSize(mask))
	fmt.Println("CIDR notation", cird)

	// Parse the CIDR and check if another IP belongs to it

	_, network, _ := net.ParseCIDR("192.168.10.0/24")
	testIP := net.ParseIP("192.168.10.50")

	if network.Contains(testIP) {
		fmt.Println("Positive", testIP, "is inside the subnet", network)
	} else {
		fmt.Println("Negative", testIP, "is NOT inside the subnet", network)
	}

}

func maskSize(mask net.IPMask) int {
	ones, _ := mask.Size()
	return ones
}
