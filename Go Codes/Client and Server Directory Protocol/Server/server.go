package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	checkError(err)
	fmt.Println("Server started on port 8080")

	for {
		conn, err := listener.Accept() // Her client baglantısı bagımsız conn nesnesidir
		if err != nil {
			continue
		}
		go handleClient(conn) // concurrent saglanir
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	cwd, _ := os.Getwd()
	reader := bufio.NewReader(conn) // Tcp'de satır satır veri okumak icin buffer saglanir

	for {
		line, err := reader.ReadString('\n') // '\n' görene kadar okuma gerceklesir
		if err != nil {
			return
		}

		line = strings.TrimSpace(line) // Bas ve sonrdaki bosluklar silinir 
		parts := strings.SplitN(line, " ", 2) // nesne iki parcaya ayrılır "cd images"
		cmd := parts[0] 

		switch cmd {
		case "dir":
			files, err := os.ReadDir(cwd) // Mevcut dzindeki dosyalar
			if err != nil {
				conn.Write([]byte("ERROR reading directory\n")) 
				continue
			}
			for _, f := range files {
				conn.Write([]byte(f.Name() + "\n")) // Her dosya ismini client'a gonderir
			}

		case "cd":
			if len(parts) < 2 {
				conn.Write([]byte("ERROR no directory specified\n"))
				continue
			}
			newDir := filepath.Join(cwd, parts[1])
			if _, err := os.Stat(newDir); os.IsNotExist(err) { // dizin var mı ? 
				conn.Write([]byte("ERROR directory not found\n"))
			} else {
				cwd = newDir
				conn.Write([]byte("OK changed to " + cwd + "\n")) 
			}

		case "pwd":
			conn.Write([]byte("Current directory: " + cwd + "\n"))

		case "quit":
			conn.Write([]byte("Bye!\n"))
			return

		default:
			conn.Write([]byte("Unknown command\n"))
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
