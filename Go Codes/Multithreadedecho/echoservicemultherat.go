// Multithreat Example
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	service := ":1201" // localhost:1201
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	chechError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	chechError(err)

	for {

		conn, err := listener.Accept() // Bir istemci baglantısı bekler ilk istek bitene kadar ikici işslem bekler
		if err != nil {
			continue
		}
		go handleClient(conn) // Her baglanti icin ayrı bir gorountine baslatır
	}
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return // Istemci baglantiyi kapattıysa cik
		}
		fmt.Println(string(buf[0:]))
		_, err = conn.Write(buf[0:n]) // Echo back
		if err != nil {
			return
		}
	}
}

func chechError(err error) {
	if err != nil {
		log.Fatalf("Error : %s", err.Error())
	}
}
