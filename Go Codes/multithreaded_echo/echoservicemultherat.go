// Package main implements a multithreaded TCP echo server.
//
// This server demonstrates concurrent connection handling in Go using goroutines.
// Each incoming client connection is handled in a separate goroutine, allowing
// the server to process multiple clients simultaneously without blocking.
//
// Architecture:
//
//	┌─────────────┐     ┌─────────────┐     ┌─────────────┐
//	│  Client 1   │────▶│             │     │ Goroutine 1 │
//	└─────────────┘     │             │────▶└─────────────┘
//	┌─────────────┐     │   Listener  │     ┌─────────────┐
//	│  Client 2   │────▶│  (Accept)   │────▶│ Goroutine 2 │
//	└─────────────┘     │             │     └─────────────┘
//	┌─────────────┐     │             │     ┌─────────────┐
//	│  Client N   │────▶│             │────▶│ Goroutine N │
//	└─────────────┘     └─────────────┘     └─────────────┘
//
// Usage:
//
//	go run echoservicemultherat.go
//
// Then connect using:
//
//	telnet localhost 1201
//
// Or using netcat:
//
//	nc localhost 1201
//
// The server echoes back any text sent to it.
package main

import (
	"fmt"
	"log"
	"net"
)

// main initializes the TCP listener and starts accepting connections.
// It binds to port 1201 on all available network interfaces.
func main() {
	// Define the service address (empty host = all interfaces)
	service := ":1201"

	// Resolve the TCP address
	// This converts the string address to a TCPAddr structure
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	// Create a TCP listener bound to the resolved address
	// This opens the port and prepares to accept connections
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	log.Printf("Multithreaded Echo Server started on %s", service)
	log.Println("Waiting for connections...")

	// Main accept loop - runs indefinitely
	for {
		// Accept blocks until a new connection arrives
		// When a client connects, Accept returns a new Conn
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}

		log.Printf("New connection from: %s", conn.RemoteAddr())

		// Spawn a new goroutine to handle this connection
		// This allows the main loop to immediately accept new connections
		// while previous clients are still being served
		go handleClient(conn)
	}
}

// handleClient processes a single client connection.
// It reads data from the connection and echoes it back until the client disconnects.
//
// Parameters:
//   - conn: The network connection to the client
//
// The function runs in its own goroutine and handles all I/O for one client.
// When the client disconnects or an error occurs, the function returns and
// the goroutine terminates.
func handleClient(conn net.Conn) {
	defer conn.Close()

	// Buffer for reading data (512 bytes is typical for simple protocols)
	var buf [512]byte

	for {
		// Read data from the client
		// n contains the number of bytes read
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Printf("Client %s disconnected", conn.RemoteAddr())
			return
		}

		// Log received data
		received := string(buf[0:n])
		fmt.Printf("Received from %s: %s", conn.RemoteAddr(), received)

		// Echo the data back to the client
		_, err = conn.Write(buf[0:n])
		if err != nil {
			log.Printf("Write error to %s: %v", conn.RemoteAddr(), err)
			return
		}
	}
}

// checkError handles fatal errors by logging and terminating the program.
// This is used for critical initialization errors where the server cannot continue.
//
// Parameters:
//   - err: The error to check
//
// If err is not nil, the program logs the error and exits with status code 1.
func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
