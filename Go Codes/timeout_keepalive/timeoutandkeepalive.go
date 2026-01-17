package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer conn.Close()

	tcpConn := conn.(*net.TCPConn)

	// 30 saniyelik genele zaman aşımı

	tcpConn.SetDeadline(time.Now().Add(30 * time.Second)) // 30 saniye boyunce mesaj gelmesse baglantıyı kapatır

	// Keepalive aktif

	tcpConn.SetKeepAlive(true)

	fmt.Println("Connected with keep-alive and timeout control.") // TCP seviyesinde baglantıyı aktif tutar
}
