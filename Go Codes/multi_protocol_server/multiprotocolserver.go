package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {

	tcpService := ":1200"
	udpService := ":1200"

	// TCP listener başlat
	go startTCPServer(tcpService)

	// UDP listener başlat
	go startUDPServer(udpService)

	for {
		time.Sleep(10 * time.Second)
	}
}

func startTCPServer(service string) {
	listener, err := net.Listen("tcp", service)
	checkError(err)
	log.Printf("TCP Server listening on %s\n", service)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("TCP Accept error:", err)
			continue
		}
		go handleTCPClient(conn)
	}
}

func handleTCPClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println("TCP Read error:", err)
			return
		}
		message := strings.TrimSpace(string(buf[0:n]))
		log.Printf("TCP Received : %s", message)

		reply := fmt.Sprintf("TCP ECHO %s", message)
		conn.Write([]byte(reply))
	}
}

func startUDPServer(service string) {
	addr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", addr)
	checkError(err)
	log.Printf("UDP Server listening on %s", service)

	var buf [512]byte
	for {
		n, clientAddr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			log.Println("UDP Read error:", err)
			continue
		}

		message := strings.TrimSpace(string(buf[0:n]))
		log.Printf("UDP received from %s: %s", clientAddr, message)

		reply := fmt.Sprintf("[UDP REPLY at %s] You said: %s", time.Now().Format(time.RFC3339), message)
		_, err = conn.WriteToUDP([]byte(reply), clientAddr)
		if err != nil {
			log.Println("UDP write error:", err)
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
