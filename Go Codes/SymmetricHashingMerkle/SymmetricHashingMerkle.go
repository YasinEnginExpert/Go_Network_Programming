package main

import (
	"bytes"
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// ======================================================
// --- 1 --- SYMMETRIC KEY ENCRYPTION (AES-128)
// ======================================================

func aesEncryptionDemo() {
	fmt.Println("===== AES Symmetric Key Demo =====")

	// AES-128 key -> Must be 16 bytes
	key := []byte("my key, len 16 b")

	// AES cipher nesnesi olusturma
	cipher, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("AES key size error", err)
		return
	}

	// Sifrelenecek veri -> tam 16 byte (AES block size)
	src := []byte("hello 16 b block")

	// AES block sifrelemek icin kullanılır
	var encrypted [16]byte
	cipher.Encrypt(encrypted[:], src)

	// Şifre çözme için buffer
	var decrypted [16]byte
	cipher.Decrypt(decrypted[:], encrypted[:])

	// Cozulen metni string olarak goster
	result := bytes.NewBuffer(nil)
	result.Write(decrypted[:])

	fmt.Println("Original:", string(src))
	fmt.Println("Encrypted (hex):", hex.EncodeToString(encrypted[:]))
	fmt.Println("Decrypted:", result.String())

}

// ======================================================
// --- 2 --- HASHING (SHA-256) — Git ve Bitcoin Bağlantılı
// ======================================================

func hashingDemo() {
	fmt.Println("\n===== Hashing (SHA-256) Demo =====")

	data := []byte("example data to hash")

	// SHA-256 hashing (Git SHA-256 geçişi ile ilgili metindeki bölüm)
	hash := sha256.Sum256(data)

	fmt.Println("Input:", string(data))
	fmt.Println("SHA-256:", hex.EncodeToString(hash[:]))
}

// ======================================================
// --- 3 --- MERKLE TREE / MERKLE ROOT ÖRNEĞİ
// 			(Git ve Bitcoin’de kullanılan yapı)
// ======================================================

// Mini Merkle Tree root hesaplama (çok sade)
func merkleRoot(values [][]byte) []byte {

	// Eğer tek eleman varsa → root budur
	if len(values) == 1 {
		h := sha256.Sum256(values[0])
		return h[:]
	}

	var nextLevel [][]byte

	// İkişer ikişer hashle
	for i := 0; i < len(values); i += 2 {

		// Eğer tek kalırsa kendini kopyalar (Bitcoin davranışı)
		if i+1 == len(values) {
			values = append(values, values[i])
		}

		combined := append(values[i], values[i+1]...)
		h := sha256.Sum256(combined)
		nextLevel = append(nextLevel, h[:])
	}

	// Rekürsif şekilde üst seviyeyi hesapla
	return merkleRoot(nextLevel)
}

func merkleDemo() {
	fmt.Println("\n===== Merkle Tree Demo =====")

	// Hashlenecek yapraklar (transaction / git object gibi düşünebilirsin)
	leaves := [][]byte{
		[]byte("tx1"),
		[]byte("tx2"),
		[]byte("tx3"),
		[]byte("tx4"),
	}

	root := merkleRoot(leaves)

	fmt.Println("Merkle Root:", hex.EncodeToString(root))
}

// ======================================================
// --- 4 --- MAIN
// ======================================================

func main() {
	aesEncryptionDemo()
	hashingDemo()
	merkleDemo()
}
