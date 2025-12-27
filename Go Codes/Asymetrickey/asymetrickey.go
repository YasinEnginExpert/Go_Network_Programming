package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {

	privatekey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publickey := &privatekey.PublicKey

	message := []byte("hello world")

	// Encrypt with public key
	ciphertext, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, publickey, message, nil)

	// Dectpyt with private key
	plaintext, _ := rsa.DecryptOAEP(sha256.New(), rand.Reader, privatekey, ciphertext, nil)

	fmt.Println("Cipher:", ciphertext)
	fmt.Println("Plain:", string(plaintext))
}
