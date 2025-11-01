// raw Socket

package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	target := "8.8.8.8" // Google DNS

	conn, err := net.ListenIP("ip4:icmp", &net.IPAddr{IP: net.ParseIP("0.0.0.0")}) // Herhangi bir arayüzden dinle
	if err != nil {
		log.Fatalf("Bağlantı oluşturulamadı %v", err)
	}
	defer conn.Close() // Main bittiğinde otomatik çagrılır

	fmt.Println("ICMP baglantısı olusturuldu, hedef:", target)

	msg := make([]byte, 8)
	msg[0] = 8  // Type (8 = Echo Request)
	msg[1] = 0  // Code
	msg[2] = 0  // Checksum (geçici)
	msg[3] = 0  // Checksum (geçici)
	msg[4] = 0  // Identifier (yüksek byte)
	msg[5] = 13 // Identifier (düşük byte)
	msg[6] = 0  // Sequence number (yüksek byte)
	msg[7] = 37 // Sequence number (düşük byte)

	// Checksum hesabi
	checksum := checksum(msg)
	msg[2] = byte(checksum >> 8)
	msg[3] = byte(checksum & 0xff)

	/*

		0       7 8      15 16             31
		+--------+--------+----------------+
		| Type   | Code   | Checksum       |
		+--------+--------+----------------+
		| Identifier       | Sequence      |
		+----------------------------------+
	*/

	dst := &net.IPAddr{IP: net.ParseIP(target)}

	start := time.Now()
	_, err = conn.WriteToIP(msg, dst)
	if err != nil {
		log.Fatalf("Ping gönderilemedi: %v", err)
	}
	fmt.Println("Ping gönderildi, yanıt bekleniyor. . .")

	// Yanıtı almak için buffer oluşturulur

	reply := make([]byte, 1500)
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	n, addr, err := conn.ReadFrom(reply)
	if err != nil {
		log.Fatalf("Yanit alinamadı: %v", err)

	}
	duration := time.Since(start)                                                            // time since starttan itibaren gecen zamanı hesaplar
	fmt.Printf("Yanıt %s adresinden geldi! %d byte, RTT = %v\n", addr.String(), n, duration) // %v : otomatik tür algılama

}

func checksum(data []byte) uint16 {
	sum := 0
	for i := 0; i < len(data)-1; i += 2 {
		sum += int(data[i])<<8 | int(data[i+1]) // | bit wise seviyesinde birleştirme işlemi
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	return uint16(^sum) // Bitwise operasyonu yapılır sayının tersini alırız
}
