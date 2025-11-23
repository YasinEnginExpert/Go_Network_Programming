package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	str := "百度一下, 你就知道 🙂"

	fmt.Println("Original string:", str)

	// Encode: string → rune slice → UTF-16 (uint16 slice)
	runes := []rune(str)
	utf16Encoded := utf16.Encode(runes)

	fmt.Println("UTF-16 uint16 slice:", utf16Encoded)

	// Decode: UTF-16 slice → runes → string
	decodedRunes := utf16.Decode(utf16Encoded)
	decoded := string(decodedRunes)

	fmt.Println("Decoded string:", decoded)
}
