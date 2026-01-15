// Package main implements an ICMP Echo Request (Ping) using raw sockets.
//
// This program demonstrates low-level network programming by constructing
// and sending an ICMP Echo Request packet manually, then receiving the
// Echo Reply to measure round-trip time (RTT).
//
// ICMP (Internet Control Message Protocol) is defined in RFC 792 and is
// used for diagnostic and error reporting purposes in IP networks.
//
// ICMP Echo Request/Reply Packet Structure (RFC 792):
//
//	 0                   1                   2                   3
//	 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|     Type      |     Code      |          Checksum             |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|           Identifier          |        Sequence Number        |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|                             Data                              |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
// Type values:
//   - 8: Echo Request (ping)
//   - 0: Echo Reply (pong)
//
// Reference: https://www.rfc-editor.org/rfc/rfc792
//
// Note: This program requires elevated privileges (Administrator on Windows,
// root on Linux/macOS) to create raw sockets.
//
// Usage:
//
//	go run ping.go
package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// Target IP address to ping
	target := "8.8.8.8" // Google's public DNS server

	// Create a raw ICMP socket
	// "ip4:icmp" specifies IPv4 with ICMP protocol
	// Binding to 0.0.0.0 allows receiving on any interface
	conn, err := net.ListenIP("ip4:icmp", &net.IPAddr{IP: net.ParseIP("0.0.0.0")})
	if err != nil {
		log.Fatalf("Failed to create ICMP connection: %v", err)
	}
	defer conn.Close()

	fmt.Printf("ICMP connection established, target: %s\n", target)

	// ============================================================
	// Construct ICMP Echo Request Packet
	// ============================================================
	// Minimum ICMP Echo packet is 8 bytes (header only, no data)
	msg := make([]byte, 8)

	// Byte 0: Type (8 = Echo Request)
	msg[0] = 8

	// Byte 1: Code (0 for Echo Request)
	msg[1] = 0

	// Bytes 2-3: Checksum (calculated after filling other fields)
	msg[2] = 0
	msg[3] = 0

	// Bytes 4-5: Identifier (used to match requests with replies)
	// Using arbitrary identifier 13 (0x000D)
	msg[4] = 0  // High byte
	msg[5] = 13 // Low byte

	// Bytes 6-7: Sequence Number (incremented per request)
	// Using arbitrary sequence 37 (0x0025)
	msg[6] = 0  // High byte
	msg[7] = 37 // Low byte

	// Calculate and set checksum
	// The checksum is the 16-bit one's complement of the one's complement
	// sum of the ICMP message starting with the Type field
	cs := calculateChecksum(msg)
	msg[2] = byte(cs >> 8)   // High byte
	msg[3] = byte(cs & 0xff) // Low byte

	// ============================================================
	// Send ICMP Echo Request
	// ============================================================
	dst := &net.IPAddr{IP: net.ParseIP(target)}

	startTime := time.Now()
	_, err = conn.WriteToIP(msg, dst)
	if err != nil {
		log.Fatalf("Failed to send ping: %v", err)
	}
	fmt.Println("Echo Request sent, waiting for reply...")

	// ============================================================
	// Receive ICMP Echo Reply
	// ============================================================
	// Buffer for receiving the reply (1500 = typical MTU size)
	reply := make([]byte, 1500)

	// Set a read deadline to avoid waiting forever
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))

	n, addr, err := conn.ReadFrom(reply)
	if err != nil {
		log.Fatalf("Failed to receive reply: %v", err)
	}

	// Calculate round-trip time
	rtt := time.Since(startTime)

	// ============================================================
	// Parse and Display Results
	// ============================================================
	fmt.Printf("\nReply received from %s\n", addr.String())
	fmt.Printf("  Bytes received: %d\n", n)
	fmt.Printf("  Round-trip time: %v\n", rtt)

	// The reply includes IP header (20 bytes) + ICMP message
	// ICMP Echo Reply has Type = 0
	if n >= 28 { // 20 byte IP header + 8 byte ICMP header
		icmpType := reply[20] // First byte of ICMP header after IP header
		icmpCode := reply[21]
		fmt.Printf("  ICMP Type: %d (0 = Echo Reply)\n", icmpType)
		fmt.Printf("  ICMP Code: %d\n", icmpCode)
	}
}

// calculateChecksum computes the ICMP checksum for the given data.
//
// The checksum algorithm is defined in RFC 792:
// "The checksum is the 16-bit one's complement of the one's complement
// sum of the ICMP message starting with the ICMP Type."
//
// Parameters:
//   - data: The ICMP message bytes to checksum
//
// Returns:
//   - The computed 16-bit checksum value
func calculateChecksum(data []byte) uint16 {
	sum := 0

	// Sum all 16-bit words
	for i := 0; i < len(data)-1; i += 2 {
		// Combine two bytes into a 16-bit word (big-endian)
		// Using bitwise OR to combine high byte (shifted left 8) with low byte
		sum += int(data[i])<<8 | int(data[i+1])
	}

	// Handle odd-length data (add last byte padded with zero)
	if len(data)%2 == 1 {
		sum += int(data[len(data)-1]) << 8
	}

	// Fold 32-bit sum to 16 bits by adding carry bits
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)

	// Take one's complement
	return uint16(^sum)
}
