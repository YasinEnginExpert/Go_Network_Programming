package main

import (
	"log"
	"net"
	"net/netip"
	"time"

	"github.com/jsimonetti/rtnetlink"
	"github.com/mdlayher/arp"
	"github.com/mdlayher/ethernet"
)

func main() {
	ifaceName := "eth0"
	vip := netip.MustParseAddr("198.51.100.6")

	// Interface bilgisi
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		log.Fatal(err)
	}

	// Netlink ile VIP ekle
	conn, err := rtnetlink.Dial(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("VIP ekleniyor:", vip)

	addr := &rtnetlink.AddressMessage{
		Family:       unixAF_INET(),
		PrefixLength: 32,
		Index:        uint32(iface.Index),
		Attributes: &rtnetlink.AddressAttributes{
			Address: vip.AsSlice(),
			Local:   vip.AsSlice(),
		},
	}

	if err := conn.Address.New(addr); err != nil {
		log.Println("VIP zaten ekli olabilir:", err)
	} else {
		log.Println("VIP eklendi:", vip)
	}

	// ARP socket (Layer 2)
	c, err := arp.Dial(iface)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Gratuitous ARP packet
	pkt, err := arp.NewPacket(
		arp.OperationReply,
		iface.HardwareAddr,
		vip,
		iface.HardwareAddr,
		vip,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Ethernet frame içine sar
	payload, err := pkt.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}

	eth := &ethernet.Frame{
		Destination: ethernet.Broadcast,
		Source:      iface.HardwareAddr,
		EtherType:   ethernet.EtherTypeARP,
		Payload:     payload,
	}

	_, err = eth.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}

	// Periyodik GARP gönder
	log.Println("Gratuitous ARP gönderiliyor...")
	for {
		if err := c.WriteTo(pkt, ethernet.Broadcast); err != nil {
			log.Println("ARP gönderme hatası:", err)
		}
		time.Sleep(2 * time.Second)
	}
}

// küçük helper (platform sadeleştirme)
func unixAF_INET() uint8 {
	return 2
}
