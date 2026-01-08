package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	server := "localhost:8080"
	path := "/lookup"

	// ÖRNEK SORGULAR
	lookupType := "mac" // mac | ip | domain
	argument := "tkng.io"

	// URL oluştur
	addr, err := url.Parse("http://" + server + path)
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{}
	params.Add(lookupType, argument)
	addr.RawQuery = params.Encode()

	// HTTP GET isteği
	resp, err := http.DefaultClient.Get(addr.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Response yazdır
	io.Copy(os.Stdout, resp.Body)
	fmt.Println()
}
