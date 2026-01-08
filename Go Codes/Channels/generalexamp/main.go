package main

import (
	"fmt"
	"sync"
	"time"
)

// Channel üzerinden taşınacak veri
type Data struct {
	Hostname string
	Version  string
	Uptime   string
}

// Worker: cihazdan bilgi alır ve channel'a gönderir
func getVersion(host string, ch chan<- Data, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simüle edilmiş network gecikmesi
	time.Sleep(time.Duration(len(host)%3+1) * time.Second)

	result := Data{
		Hostname: host,
		Version:  "simulated-version",
		Uptime:   "simulated-uptime",
	}

	// Veriyi channel üzerinden gönder
	ch <- result
}

// Consumer: channel'dan gelen verileri yazdırır
func printer(in <-chan Data) {
	for data := range in {
		fmt.Printf(
			"Hostname: %s\nVersion: %s\nUptime: %s\n\n",
			data.Hostname,
			data.Version,
			data.Uptime,
		)
	}
}

func main() {
	hosts := []string{
		"sandbox-iosxe",
		"sandbox-nxos",
		"sandbox-iosxr",
	}

	ch := make(chan Data)

	// Yazdırıcı goroutine
	go printer(ch)

	var wg sync.WaitGroup

	// Her cihaz için bir goroutine
	for _, h := range hosts {
		wg.Add(1)
		go getVersion(h, ch, &wg)
	}

	// Tüm worker'lar bitince channel'ı kapat
	wg.Wait()
	close(ch)
}


