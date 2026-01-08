package main

import (
	"fmt"
	"net"

	"github.com/yl2chen/cidranger"
)

func main() {

	// Path-compressed trie oluştur
	ranger := cidranger.NewPCTrieRanger()

	// CIDR listesi
	IPs := []string{
		"100.64.0.0/16",
		"127.0.0.0/8",
		"172.16.0.0/16",
		"192.0.2.0/24",
		"192.0.2.0/25",
		"192.0.2.127/25",
	}

	// Trie içine ekle
	for _, prefix := range IPs {
		_, network, err := net.ParseCIDR(prefix)
		if err != nil {
			panic(err)
		}
		ranger.Insert(
			cidranger.NewBasicRangerEntry(*network),
		)
	}

	// ---------------------------------------------------
	// 1) Belirli bir IP herhangi bir prefix içinde mi?
	// ---------------------------------------------------
	checkIP := "127.0.0.1"

	ok, err := ranger.Contains(net.ParseIP(checkIP))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Does the range contain %s?: %v\n", checkIP, ok)

	// ---------------------------------------------------
	// 2) Bir IP'yi kapsayan TÜM network'leri bul
	// ---------------------------------------------------
	netIP := "192.0.2.18"

	nets, err := ranger.ContainingNetworks(net.ParseIP(netIP))
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nNetworks that contain IP address %s ->\n", netIP)
	for _, e := range nets {
		fmt.Println(e.Network())
	}
}
