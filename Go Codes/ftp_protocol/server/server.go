// Package main implements a simple FTP-like file server using text-based protocol.
//
// This server provides basic file system navigation commands similar to FTP:
//   - DIR: List directory contents
//   - CD <path>: Change current directory
//   - PWD: Print working directory
//
// Protocol Specification:
//
//	Client sends: COMMAND [ARGUMENT]\r\n
//	Server responds: RESULT\r\n (or multiple lines for DIR)
//
// Example Session:
//
//	Client: PWD
//	Server: C:\Users\example
//	Client: DIR
//	Server: file1.txt
//	        file2.go
//	        subdir
//
//	Client: CD subdir
//	Server: OK
//
// This is an educational example demonstrating:
//   - Text-based protocol parsing
//   - Concurrent client handling with goroutines
//   - File system operations in Go
//
// Usage:
//
//	go run server.go
//
// Connect using:
//
//	telnet localhost 1202
//	nc localhost 1202
package main

import (
	"log"
	"net"
	"os"
	"strings"
)

// Protocol commands
const (
	DIR = "DIR" // List directory contents
	CD  = "CD"  // Change directory
	PWD = "PWD" // Print working directory
)

func main() {
	// Bind to all interfaces on port 1202
	service := "0.0.0.0:1202"

	// Resolve the TCP address
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	// Create TCP listener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	log.Printf("FTP Server started on %s", service)
	log.Println("Available commands: DIR, CD <path>, PWD")

	// Main accept loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}

		log.Printf("New connection from %s", conn.RemoteAddr())

		// Handle each client in a separate goroutine
		go handleClient(conn)
	}
}

// handleClient processes commands from a single client connection.
// It reads commands, parses them, and dispatches to appropriate handlers.
//
// Parameters:
//   - conn: The client connection
//
// The function runs until the client disconnects or an error occurs.
func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte

	for {
		// Read command from client
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Printf("Client %s disconnected", conn.RemoteAddr())
			return
		}

		// Parse command: split into command and arguments
		input := strings.TrimSpace(string(buf[0:n]))
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := strings.ToUpper(parts[0])
		log.Printf("Received from %s: %s", conn.RemoteAddr(), input)

		// Dispatch command to appropriate handler
		switch command {
		case CD:
			if len(parts) < 2 {
				sendResponse(conn, "ERROR: CD requires a path argument")
			} else {
				changeDirectory(conn, parts[1])
			}

		case DIR:
			listDirectory(conn)

		case PWD:
			printWorkingDirectory(conn)

		default:
			log.Printf("Unknown command from %s: %s", conn.RemoteAddr(), command)
			sendResponse(conn, "ERROR: Unknown command. Use DIR, CD, or PWD")
		}
	}
}

// changeDirectory changes the server's current working directory.
//
// Parameters:
//   - conn: Client connection for response
//   - path: Target directory path
//
// Responds with "OK" on success or "ERROR" on failure.
func changeDirectory(conn net.Conn, path string) {
	err := os.Chdir(path)
	if err != nil {
		log.Printf("CD error: %v", err)
		sendResponse(conn, "ERROR: "+err.Error())
		return
	}
	sendResponse(conn, "OK")
}

// printWorkingDirectory sends the current working directory to the client.
//
// Parameters:
//   - conn: Client connection for response
func printWorkingDirectory(conn net.Conn) {
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("PWD error: %v", err)
		sendResponse(conn, "ERROR: "+err.Error())
		return
	}
	sendResponse(conn, wd)
}

// listDirectory sends the contents of the current directory to the client.
// Each entry is sent on a separate line, followed by an empty line.
//
// Parameters:
//   - conn: Client connection for response
func listDirectory(conn net.Conn) {
	// Open current directory
	dir, err := os.Open(".")
	if err != nil {
		log.Printf("DIR error: %v", err)
		sendResponse(conn, "ERROR: "+err.Error())
		return
	}
	defer dir.Close()

	// Read all directory entries
	entries, err := dir.Readdirnames(-1)
	if err != nil {
		log.Printf("Readdirnames error: %v", err)
		sendResponse(conn, "ERROR: "+err.Error())
		return
	}

	// Send each entry on its own line
	for _, name := range entries {
		conn.Write([]byte(name + "\r\n"))
	}

	// Send empty line to indicate end of listing
	conn.Write([]byte("\r\n"))
}

// sendResponse sends a text response to the client with CRLF line ending.
//
// Parameters:
//   - conn: Client connection
//   - msg: Message to send
func sendResponse(conn net.Conn, msg string) {
	conn.Write([]byte(msg + "\r\n"))
}

// checkError handles fatal errors during server initialization.
// If an error occurs, it logs the error and terminates the program.
//
// Parameters:
//   - err: The error to check
func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
