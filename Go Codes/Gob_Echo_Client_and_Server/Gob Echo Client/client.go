package main

import (
	"encoding/gob"
	"log"
	"net"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "host:port")
	}

	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkerror(err)
	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	person := Person{
		Name: Name{Family: "Yasins", Personal: "Yasin"},
		Email: []Email{
			{Kind: "home", Address: "Yasin@Yasins.name"},
			{Kind: "work", Address: "Yasin.Hacıh@yasins.edu.tr"},
		},
	}

	for n := 0; n < 10; n++ {
		encoder.Encode(person) // gönder
		var newPerson Person
		decoder.Decode(&newPerson) // al
		log.Println("Client received:", newPerson)
	}
}

func checkerror(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
