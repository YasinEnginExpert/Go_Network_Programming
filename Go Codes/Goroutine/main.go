package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

type Router struct {
	Hostname string `yaml:"hostname"`
	Platform string `yaml:"platform"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Inventory struct {
	Routers []Router `yaml:"router"`
}

// Simulated network version check
func getVersion(r Router) {
	// Platforma göre farklı gecikme süreleri simüle et
	// Bu, gerçek bir ağ isteği yerine geçer
	// Örneğin, Cisco IOS XE için 2 saniye, NX-OS için 3 saniye, IOS XR için 1 saniye
	switch r.Platform {
	case "cisco_iosxe":
		time.Sleep(2 * time.Second)
	case "cisco_nxos":
		time.Sleep(3 * time.Second)
	case "cisco_iosxr":
		time.Sleep(1 * time.Second)
	}

	// Sonucu yazdır
	fmt.Println("----- Router Info -----")
	fmt.Println("Hostname :", r.Hostname)
	fmt.Println("Platform :", r.Platform)
	fmt.Println("Username :", r.Username)
	fmt.Println("Version  : simulated")
	fmt.Println("-----------------------------")
}

func main() {
	// YAML dosyasını aç
	src, err := os.Open("input.yml")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	// Decode et
	decoder := yaml.NewDecoder(src)
	var inv Inventory
	if err := decoder.Decode(&inv); err != nil {
		panic(err)
	}

	start := time.Now()

	var wg sync.WaitGroup

	for _, r := range inv.Routers {
		wg.Add(1)

		go func(router Router) {
			defer wg.Done()
			getVersion(router)
		}(r)
	}

	wg.Wait()

	fmt.Println("Total execution time:", time.Since(start))
}
