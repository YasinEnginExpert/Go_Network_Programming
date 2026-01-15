// Package main implements a Daytime Protocol server as defined in RFC 867.
//
// The Daytime Protocol is a simple network service that returns the current
// date and time as a human-readable string. This implementation uses TCP
// on port 1200.
//
// Protocol Specification (RFC 867):
//   - Server listens on a well-known port (13 standardly, 1200 here)
//   - Client connects to the server
//   - Server immediately sends the current time as ASCII text
//   - Server closes the connection
//   - No request is required from the client
//
// Reference: https://www.rfc-editor.org/rfc/rfc867
//
// Usage:
//
//	go run daytimeserver.go
//
// Test with telnet:
//
//	telnet localhost 1200
//
// Or with netcat:
//
//	nc localhost 1200
//
// Note: On Windows, you may need to enable the Telnet Client feature:
//
//	Control Panel → Programs → Turn Windows features on or off → Telnet Client
package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// main starts the Daytime Protocol server.
// It binds to TCP port 1200 and serves the current time to each connecting client.
func main() {
	// Service address - port 1200 on all interfaces
	// Standard daytime port is 13, but we use 1200 to avoid requiring root privileges
	service := ":1200"

	// Resolve the TCP address from the service string
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	// Create a TCP listener
	// This opens the port and prepares to accept incoming connections
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	log.Printf("Daytime Server started on %s (RFC 867)", service)
	log.Println("Waiting for connections...")

	// Main server loop - accept and handle connections forever
	for {
		// Accept blocks until a client connects
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}

		// Get the current time as a formatted string
		// RFC 867 specifies ASCII text but doesn't mandate a specific format
		daytime := time.Now().Format(time.RFC1123)

		// Send the time to the client
		_, err = conn.Write([]byte(daytime + "\n"))
		if err != nil {
			log.Printf("Write error: %v", err)
		}

		// Log the transaction
		fmt.Printf("Served time to %s: %s\n", conn.RemoteAddr(), daytime)

		// Close the connection immediately (per RFC 867 specification)
		conn.Close()
	}
}

// checkError handles fatal errors during server initialization.
// If an error occurs, it logs the error message and terminates the program.
//
// Parameters:
//   - err: The error to check
func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
