package main

import (
	"context"
	"log"
	"net"

	pb "simplegrpcserver/proto/gen" // Derlenmiş protobuf dosyaları
	farewellpd "simplegrpcserver/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// ----------------------------------------------------
//
//	gRPC Server Struct
//	Hem Calculator hem Greeter servislerini implemente eder
//
// ----------------------------------------------------
type server struct {
	pb.UnimplementedCalculatorServer // Interface'in future-proof olması için zorunlu
	pb.UnimplementedGreeterServer
	farewellpd.UnimplementedAufWiedersehenServer
}

// ----------------------------------------------------
//
//	Calculator.Add RPC IMPLEMENTATION
//
// ----------------------------------------------------
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {

	// İş mantığı: iki sayıyı topla
	sum := req.A + req.B
	log.Println("Add() called — Sum:", sum)

	// Yanıtı oluşturup geri döndür
	return &pb.AddResponse{
		Sum: sum,
	}, nil
}

// ----------------------------------------------------
//
//	Greeter.Greet RPC IMPLEMENTATION
//
// ----------------------------------------------------
func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	// Basit bir mesaj oluştur
	msg := "Hello, " + req.Name + "!"
	log.Println("Greet() called — Message:", msg)

	return &pb.HelloResponse{
		Message: msg,
	}, nil
}

// ----------------------------------------------------
//
//	AufWiedersehen.BigGoodBye RPC IMPLEMENTATION
//
// ----------------------------------------------------
func (s *server) BigGoodBye(ctx context.Context, req *farewellpd.GoodByeRequest) (*farewellpd.GoodByeResponse, error) {

	// Basit bir mesaj oluştur
	msg := "Goodbye, " + req.Name + "!"
	log.Println("Goodbye() called — Message:", msg)

	return &farewellpd.GoodByeResponse{
		Message: msg,
	}, nil
}

func main() {

	// Sertifika dosyaları
	cert := "cert.pem"
	key := "key.pem"

	port := "50051"

	// ----------------------------------------------------
	// 1. TCP portunu dinlemeye başla
	// ----------------------------------------------------
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	// ----------------------------------------------------
	// 2. TLS sertifikalarını yükle (Server Side)
	// ----------------------------------------------------
	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalln("Failed to load TLS certificates:", err)
	}

	// ----------------------------------------------------
	// 3. Güvenli bir gRPC server oluştur
	// ----------------------------------------------------
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	// ----------------------------------------------------
	// 4. Calculator ve Greeter servislerini register et
	// ----------------------------------------------------
	pb.RegisterCalculatorServer(grpcServer, &server{})
	pb.RegisterGreeterServer(grpcServer, &server{})
	farewellpd.RegisterAufWiedersehenServer(grpcServer, &server{})

	log.Println("Secure gRPC Server running on port:", port)

	// ----------------------------------------------------
	// 5. Server'ı başlat
	// ----------------------------------------------------
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
