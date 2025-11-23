package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started on :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	state := "login"
	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		line = strings.TrimSpace(line)

		switch state {

		case "login":
			parts := strings.Split(line, " ")
			if len(parts) == 3 && parts[0] == "LOGIN" {
				user := parts[1]
				pass := parts[2]

				if user == "admin" && pass == "123" {
					conn.Write([]byte("OK\n"))
					state = "file_transfer"
				} else {
					conn.Write([]byte("FAILED\n"))
				}
			} else {
				conn.Write([]byte("FAILED\n"))
			}

		case "file_transfer":

			if line == "quit" {
				conn.Write([]byte("bye\n"))
				return
			}

			if strings.HasPrefix(line, "DIR") {
				files, _ := os.ReadDir(".")
				for _, f := range files {
					conn.Write([]byte(f.Name() + "\n"))
				}
				conn.Write([]byte("\n"))
			}

			if strings.HasPrefix(line, "GET ") {
				filename := strings.TrimPrefix(line, "GET ")
				data, err := os.ReadFile(filename)
				if err != nil {
					conn.Write([]byte("FAILED\n"))
				} else {
					conn.Write([]byte(string(data) + "\n"))
				}
			}

			if strings.HasPrefix(line, "CD ") {
				dir := strings.TrimPrefix(line, "CD ")
				if os.Chdir(dir) == nil {
					conn.Write([]byte("SUCCEEDED\n"))
				} else {
					conn.Write([]byte("FAILED\n"))
				}
			}
		}
	}
}
