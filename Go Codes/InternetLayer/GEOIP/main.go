package main

import (
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
)

func main() {

	db, err := geoip2.Open("GeoIP2-City-Test.mmdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	IPs := []string{
		"81.2.69.143",
	}

	fmt.Println("Find information for each IP:")

	for _, ipStr := range IPs {
		ip := net.ParseIP(ipStr)

		record, err := db.City(ip)
		if err != nil {
			panic(err)
		}

		fmt.Printf("\nAddress: %s\n", ipStr)
		fmt.Printf("City name: %v\n", record.City.Names["en"])
		fmt.Printf("Country name: %v\n", record.Country.Names["en"])
		fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
		fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
		fmt.Printf(
			"Coordinates: %v, %v\n",
			record.Location.Latitude,
			record.Location.Longitude,
		)
	}
}
