package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"unicode/utf16"
)

func readUint16(conn net.Conn, order binary.ByteOrder) (uint16, error) {
	var buf [2]byte
	_, err := conn.Read(buf[:])
	if err != nil {
		return 0, err
	}
	return order.Uint16(buf[:]), nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: utf16client host:port")
	}

	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 1) BOM oku
	var bomBuf [2]byte
	_, err = conn.Read(bomBuf[:])
	if err != nil {
		log.Fatal("Cannot read BOM:", err)
	}

	var order binary.ByteOrder

	bom := binary.BigEndian.Uint16(bomBuf[:])
	switch bom {
	case 0xFEFF:
		order = binary.BigEndian
		fmt.Println("Detected BOM: UTF-16BE")
	case 0xFFFE:
		order = binary.LittleEndian
		fmt.Println("Detected BOM: UTF-16LE")
	default:
		log.Fatal("Unknown BOM:", bom)
	}

	// 2) Tüm veriyi UTF-16 buffer’a oku
	var utf16Units []uint16
	for {
		u, err := readUint16(conn, order)
		if err != nil {
			break
		}
		utf16Units = append(utf16Units, u)
	}

	// 3) Decode UTF-16 → string
	runes := utf16.Decode(utf16Units)
	str := string(runes)

	fmt.Println("Received text:")
	fmt.Println(str)
}
