package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		fmt.Print("Command: ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		// Send command
		conn.Write([]byte(cmd + "\n"))

		if cmd == "quit" {
			fmt.Println("Disconnected.")
			return
		}

		if strings.HasPrefix(cmd, "DIR") {
			// DIR returns multiple lines ending with empty line
			for {
				resp, _ := serverReader.ReadString('\n')
				resp = strings.TrimSpace(resp)
				if resp == "" {
					break
				}
				fmt.Println("Server:", resp)
			}
			continue
		}

		// All other commands → read ONE line response
		resp, _ := serverReader.ReadString('\n')
		resp = strings.TrimSpace(resp)
		fmt.Println("Server:", resp)
	}
}
