// IPv4Router.go
package main

import (
	"fmt"
	"net"
)

func main() {
	// Define routing table

	routingTable := []struct {
		subnetmask net.IP
		network    net.IP
		nextHop    net.IP
	}{
		{net.IP{255, 255, 255, 240}, net.IP{192, 17, 7, 208}, net.IP{192, 12, 7, 15}},
		{net.IP{255, 255, 255, 240}, net.IP{192, 17, 7, 144}, net.IP{192, 12, 7, 67}},
		{net.IP{255, 255, 255, 0}, net.IP{192, 17, 7, 0}, net.IP{192, 12, 7, 251}},
		{net.IP{0, 0, 0, 0}, net.IP{0, 0, 0, 0}, net.IP{10, 10, 10, 10}}, // Default route catch all
	}

	// Define packets to be routed
	packets := []struct {
		sourceAddr      net.IP
		destinationAddr net.IP
		data            string
	}{
		{net.IP{1, 2, 3, 4}, net.IP{2, 3, 4, 5}, "Unknown destination"},
		{net.IP{192, 17, 7, 20}, net.IP{192, 17, 7, 251}, "Better be local"},
	}

	// Simulate routing logic
	for _, packet := range packets {
		routed := false // Tabloda eşleşme yok
		for _, entry := range routingTable {
			// Apply subnet mask to destination
			maskedDest := packet.destinationAddr.Mask(net.IPMask(entry.subnetmask))
			if maskedDest.Equal(entry.network) {
				fmt.Printf("For destination %s next hop is %s\n", packet.destinationAddr, entry.nextHop)
				routed = true
				break
			}
		}
		if !routed {
			// fallback to default route
			fmt.Printf("For destination %s next hop is default route\n", packet.destinationAddr)
		}
	}
}
