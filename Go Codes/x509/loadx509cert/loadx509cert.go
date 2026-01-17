package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {

	// ---- CERTIFICATE READ ----
	derBytes, err := os.ReadFile("jan.newmarch.name.cer")
	checkError(err)

	cert, err := x509.ParseCertificate(derBytes)
	checkError(err)

	fmt.Printf("Name: %s\n", cert.Subject.CommonName)
	fmt.Printf("Not before: %s\n", cert.NotBefore)
	fmt.Printf("Not after:  %s\n", cert.NotAfter)

	// ---- LOAD PUBLIC KEY ----
	pubFile, err := os.Open("public.key")
	checkError(err)
	defer pubFile.Close()

	dec := gob.NewDecoder(pubFile)
	pubKey := new(rsa.PublicKey)
	checkError(dec.Decode(pubKey))

	// ---- COMPARE PUBLIC KEYS ----
	certKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		log.Fatalln("Certificate does not contain RSA public key")
	}

	if certKey.E == pubKey.E && certKey.N.Cmp(pubKey.N) == 0 {
		fmt.Println("Same public key")
	} else {
		fmt.Println("Different public key")
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error:", err)
	}
}
