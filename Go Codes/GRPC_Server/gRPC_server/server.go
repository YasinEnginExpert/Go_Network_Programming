package main

import (
	"context"
	"log"
	"net"

	pb "simplegrpcserver/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

// Add RPC implementation
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	log.Println("Sum", sum)
	return &pb.AddResponse{
		Sum: sum,
	}, nil
}

func main() {

	cert := "cert.pem"
	key := "key.pem"
	port := "50051"

	// Dinleme adresi mutlaka ":port" formatında olmalı
	list, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	creads, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalln("Failed to load credentials", err)
	}

	grpcserver := grpc.NewServer(grpc.Creds(creads))

	// Servisi gRPC server'a register ediyoruz
	pb.RegisterCalculatorServer(grpcserver, &server{})

	log.Println("Server running on port:", port)

	if err := grpcserver.Serve(list); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
