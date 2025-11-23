package main

import (
	"fmt"
	"log"
	"net/textproto"
)

func check(err error) {
	if err != nil {
		log.Fatalln("error:", err)
	}
}

func main() {

	// Connect to a Unix-domain socket acting as gake webserver
	conn, err := textproto.Dial("unix", "/tmp/fakewebserver")
	check(err)
	defer conn.Close()

	fmt.Println("Sending GET request to /mypage")

	// Send a command like GET /mypage
	// Cmd returns an ID so we can pair request/response together
	id, err := conn.Cmd("GET /mypage")
	check(err)

	// Tell textproto that we expect a response for this ID
	conn.StartResponse(id)
	defer conn.EndResponse(id)

	// Read a line that MUST start with status code 200
	status, message, err := conn.ReadCodeLine(200)
	check(err)

	fmt.Println("Status Code:", status)
	fmt.Println("Message:", message)
	fmt.Println("No error:", err)
	
}

/*
Bir terminalde nc -lkU /tmp/fakewebserver çalıştır bu netcat TCP/UDP/unix-domain socket ile elle mesaj göndermeye yarayan bir araçtır.
Netcat ile sahte bir sunucu gibi davranıyor 

*/

