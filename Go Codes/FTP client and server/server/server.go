/* FTP Server - Simple Text Protocol Example */

package main

import (
	"log"
	"net"
	"os"
	"strings"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	service := "0.0.0.0:1202"

	// TCP adresini çözümle
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	// TCP dinleyici oluştur
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	log.Println("FTP Server started on", service)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			conn.Close()
			return
		}

		// Komutu ayrıştır
		s := strings.Split(strings.TrimSpace(string(buf[0:n])), " ")
		log.Println("Received:", s)

		if len(s) == 0 {
			continue
		}

		switch s[0] {
		case CD:
			if len(s) < 2 {
				conn.Write([]byte("ERROR"))
			} else {
				chdir(conn, s[1])
			}
		case DIR:
			dirList(conn)
		case PWD:
			pwd(conn)
		default:
			log.Println("Unknown command:", s)
			conn.Write([]byte("ERROR"))
		}
	}
}

func chdir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))
}

func dirList(conn net.Conn) {
	defer conn.Write([]byte("\r\n")) // boş satırla bitir
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error:", err.Error())
	}
}
