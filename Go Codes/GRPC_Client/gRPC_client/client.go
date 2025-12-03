package main

import (
	"context"
	"log"
	mainapipb "simplegrpcserver/proto/gen"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	// Sertifikayı yükle
	creds, err := credentials.NewClientTLSFromFile("cert.pem", "localhost")
	if err != nil {
		log.Fatalln("Failed to load TLS certificate:", err)
	}

	// TLS ile bağlan
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer conn.Close()

	client := mainapipb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := mainapipb.AddRequest{
		A: 10,
		B: 20,
	}

	res, err := client.Add(ctx, &req)
	if err != nil {
		log.Fatalln("Add RPC failed:", err)
	}

	log.Println("Sum:", res.Sum)
}
