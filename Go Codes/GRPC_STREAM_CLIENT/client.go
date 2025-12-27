package main

import (
	"context"
	"io"
	"log"

	mainpb "grpcstreamclient/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// ---Sunucuya bağlan---
	conn, err := grpc.Dial("localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer conn.Close()

	//---Client instance oluştur---
	client := mainpb.NewCalculatorClient(conn)

	//Request oluştur---
	req := &mainpb.FibonacciRequest{
		N: 10,
	}

	//---context oluştur---
	ctx := context.Background()

	// ---Streaming RPC'i çağır---
	stream, err := client.GenerateFibonacci(ctx, req)
	if err != nil {
		log.Fatalln("Error calling GenerateFibonacci:", err)
	}

	// ---Server'dan gelen stream'i oku---
	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			log.Println(">> Stream finished")
			break
		}

		if err != nil {
			log.Fatalln("Error receiving stream:", err)
		}

		log.Println("Fibonacci number:", resp.GetNumber())
	}
}
