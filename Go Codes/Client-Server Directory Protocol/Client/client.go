package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // TCP baglantisi olusturulur
	checkError(err)
	defer conn.Close()

	fmt.Println("Connected to directory server.")
	reader := bufio.NewReader(os.Stdin) // Kullanıcıdan terminal girisi
	serverReader := bufio.NewReader(conn) // Sunucudan gelen yanıtları okumak icin

	for {
		fmt.Print("> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		// Send command to server
		conn.Write([]byte(cmd + "\n"))

		if cmd == "quit" {
			break
		}

		// Read response
		for {
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 200)) // dongu kirici
			line, err := serverReader.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Print(line) // her satir alinir ve yazilir
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
