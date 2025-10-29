package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s network-type service\n", os.Args[0])
	}

	networkType := os.Args[1] // tcp veya udp
	service := os.Args[2]     // servis adı

	port, err := net.LookupPort(networkType, service)
	if err != nil {
		log.Fatalln("Error :", err)
	}

	fmt.Println("Serive port", port)

}
