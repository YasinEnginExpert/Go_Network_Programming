package main

import (
	"encoding/binary"
	"log"
	"net"
	"unicode/utf16"
)

const (
	BOM_BE = 0xFEFF // big-endian
	BOM_LE = 0xFFFE // little-endian
)

func main() {
	listener, err := net.Listen("tcp", ":1210")
	if err != nil {
		log.Fatal("Cannot listen:", err)
	}
	log.Println("UTF-16 server running on :1210")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// Mesaj
	msg := "UTF-16 服务器发送的文本: 百度一下, 你就知道 🙂"

	// UTF-16 encode
	utf16Units := utf16.Encode([]rune(msg))

	// ▣ SERVER ALWAYS SENDS BIG ENDIAN (clean standard)
	// İlk 2 byte → BOM (FE FF)
	binary.Write(conn, binary.BigEndian, uint16(BOM_BE))

	// UTF-16 kod birimlerini gönder
	for _, u := range utf16Units {
		binary.Write(conn, binary.BigEndian, u)
	}
}
