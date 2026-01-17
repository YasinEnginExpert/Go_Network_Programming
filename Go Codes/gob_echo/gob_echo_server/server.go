package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personla string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	service := ":1200" // TCP Port
	tcpAddre, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddre)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		for n := 0; n < 10; n++ {
			var person Person
			decoder.Decode(&person) // Istemciden okurum
			fmt.Println("Server received:", person)
			encoder.Encode(person) // Geri fırlat
		}
		conn.Close()
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
