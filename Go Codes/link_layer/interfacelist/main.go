package main

import (
	"log"

	"github.com/jsimonetti/rtnetlink/rtnl"
)

func main() {
	conn, err := rtnl.Dial(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	links, err := conn.Links()
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range links {
		log.Printf("Interface: %s | Flags: %v\n", l.Name, l.Flags)
	}

}
