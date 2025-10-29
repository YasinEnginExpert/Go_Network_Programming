package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Windows ciahzında telnet localhost 1200 yap çalışmassa windows özellikelrden telnet client özelligini ac

func main() {

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // Istemciye saati gönderir
		fmt.Println(daytime)
		conn.Close() // baglantı kapatılır
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
