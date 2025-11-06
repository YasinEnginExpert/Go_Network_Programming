package main

import (
	"encoding/asn1"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	service := ":1200" // Port numaramız

	tcpAddr, err := net.ResolveTCPAddr("tcp", service) // TCP adresini çözümleyeceğim
	fmt.Println(tcpAddr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr) // TCP sunucusunu baslat
	checkError(err)

	log.Println("Server listening on", service)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now()
		mdata, _ := asn1.Marshal(daytime)

		conn.Write(mdata)
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error:", err)
	}
}
