package main

import (
	"fmt"
	"log"
	"net/http"
)

// /check endpoint'i → health check
func checkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

// /lookup endpoint'i → query bazlı işlem
func lookupHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Domain lookup
	if domain := query.Get("domain"); domain != "" {
		fmt.Fprintf(w,
			"Domain Name: %s\nRegistrar: Example Registrar\n",
			domain,
		)
		return
	}

	// IP lookup
	if ip := query.Get("ip"); ip != "" {
		fmt.Fprintf(w,
			"IP Address: %s\nOwner: Example ISP\n",
			ip,
		)
		return
	}

	// MAC lookup
	if mac := query.Get("mac"); mac != "" {
		fmt.Fprintf(w,
			"MAC Address: %s\nVendor: Example Vendor\n",
			mac,
		)
		return
	}

	// Hatalı istek
	http.Error(w, "Invalid lookup type", http.StatusBadRequest)
}

func main() {
	// Handler'ları DefaultServeMux'e kaydet
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/lookup", lookupHandler)

	// HTTP Server tanımı
	srv := http.Server{
		Addr: "0.0.0.0:8080",
	}

	log.Println("HTTP server listening on 0.0.0.0:8080")

	// Server başlat
	log.Fatal(srv.ListenAndServe())
}
