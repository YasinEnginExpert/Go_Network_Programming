package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	random := rand.Reader
	var key rsa.PrivateKey

	// Load RSA private key
	loadKey("private.key", &key)

	now := time.Now()
	// ✔ Correct 1-year validity
	then := now.Add(365 * 24 * time.Hour)

	// Certificate Template
	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),

		Subject: pkix.Name{
			CommonName:   "jan.newmarch.name",
			Organization: []string{"Jan Newmarch"},
		},

		NotBefore: now,
		NotAfter:  then,

		SubjectKeyId: []byte{1, 2, 3, 4},

		KeyUsage: x509.KeyUsageCertSign |
			x509.KeyUsageKeyEncipherment |
			x509.KeyUsageDigitalSignature,

		BasicConstraintsValid: true,
		IsCA:                  true,

		DNSNames: []string{"jan.newmarch.name", "localhost"},
	}

	// Self-signed certificate (template signs itself)
	derBytes, err := x509.CreateCertificate(random, &template, &template, &key.PublicKey, &key)
	checkError(err)

	// Save .cer (DER format)
	certCerFile, err := os.Create("jan.newmarch.name.cer")
	checkError(err)
	_, _ = certCerFile.Write(derBytes)
	certCerFile.Close()

	// Save .pem (PEM format)
	certPEMFile, err := os.Create("jan.newmarch.name.pem")
	checkError(err)
	pem.Encode(certPEMFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	})
	certPEMFile.Close()

	// Save private key as PEM
	keyPEMFile, err := os.Create("private.pem")
	checkError(err)
	pem.Encode(keyPEMFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(&key),
	})
	keyPEMFile.Close()

	fmt.Println("Certificate and key successfully generated!")
}

// Load RSA private key stored with gob
func loadKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	defer inFile.Close()

	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}
