package main

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// go run asn_daytime_client.go localhost:1200
func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "host:port")
	}

	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	var newTime time.Time
	_, err = asn1.Unmarshal(result, &newTime)
	checkError(err)

	fmt.Println("After marshal/unmarshal:", newTime.String())
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error:", err)
	}
}
