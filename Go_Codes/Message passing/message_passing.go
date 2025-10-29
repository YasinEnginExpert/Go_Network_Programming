package main

import (
	"fmt"
	"time"
)

func worker(messages chan string) { // Message değişkeni üzerinden string mesajları alabilir veya gönderebilrim

	for msg := range messages {
		fmt.Println("Received:", msg)
	}
}

func main() {
	messages := make(chan string) // Yeni kanal oluştudum bu kanal string türünde veri taşır | goroutine’ler arası veri boru hattı
	go worker(messages)           // go parametresiyle fonksiyon ayrı bir goroutine oalrak çalışmaya başlar arka planda eşzamanlı çalışır

	messages <- "Hello" // <- Knala veri göndermek için kullanılır
	messages <- "from"
	messages <- "Go!"

	close(messages) // Kanalımızı kapatırız

	time.Sleep(time.Second)
}
