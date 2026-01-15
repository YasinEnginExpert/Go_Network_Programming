// Package main implements a secure gRPC server with multiple services.
//
// This server demonstrates:
//   - Multiple gRPC service registration on a single server
//   - TLS/SSL encryption for secure communication
//   - Unary RPC implementations
//
// Services Implemented:
//  1. Calculator: Provides arithmetic operations (Add)
//  2. Greeter: Provides greeting messages (Greet)
//  3. AufWiedersehen: Provides farewell messages (BigGoodBye)
//
// gRPC Architecture:
//
//	┌─────────────┐         TLS          ┌─────────────────────────┐
//	│   Client    │◄──────────────────►  │      gRPC Server        │
//	│             │    Port 50051        │  ┌─────────────────────┐│
//	│  - Stub     │                      │  │  CalculatorService  ││
//	│  - Channel  │                      │  │  GreeterService     ││
//	│             │                      │  │  FarewellService    ││
//	└─────────────┘                      │  └─────────────────────┘│
//	                                     └─────────────────────────┘
//
// TLS Configuration:
//   - Requires cert.pem (certificate) and key.pem (private key)
//   - Generate self-signed certificates for testing:
//     openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
//
// Usage:
//
//	go run server.go
//
// The server listens on port 50051 with TLS encryption.
package main

import (
	"context"
	"log"
	"net"

	pb "simplegrpcserver/proto/gen"
	farewellpd "simplegrpcserver/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// server implements all gRPC service interfaces.
// It embeds the Unimplemented* types for forward compatibility.
// This ensures the server compiles even if new methods are added to the proto.
type server struct {
	pb.UnimplementedCalculatorServer
	pb.UnimplementedGreeterServer
	farewellpd.UnimplementedAufWiedersehenServer
}

// Add implements the Calculator.Add RPC method.
// It receives two integers and returns their sum.
//
// Parameters:
//   - ctx: Context for the RPC call (can contain deadlines, cancellation signals)
//   - req: AddRequest containing the two numbers to add
//
// Returns:
//   - AddResponse containing the sum
//   - error if the operation fails
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	log.Printf("Add() called: %d + %d = %d", req.A, req.B, sum)

	return &pb.AddResponse{
		Sum: sum,
	}, nil
}

// Greet implements the Greeter.Greet RPC method.
// It receives a name and returns a personalized greeting.
//
// Parameters:
//   - ctx: Context for the RPC call
//   - req: HelloRequest containing the name to greet
//
// Returns:
//   - HelloResponse containing the greeting message
//   - error if the operation fails
func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	msg := "Hello, " + req.Name + "!"
	log.Printf("Greet() called: %s", msg)

	return &pb.HelloResponse{
		Message: msg,
	}, nil
}

// BigGoodBye implements the AufWiedersehen.BigGoodBye RPC method.
// It receives a name and returns a personalized farewell message.
//
// Parameters:
//   - ctx: Context for the RPC call
//   - req: GoodByeRequest containing the name to bid farewell
//
// Returns:
//   - GoodByeResponse containing the farewell message
//   - error if the operation fails
func (s *server) BigGoodBye(ctx context.Context, req *farewellpd.GoodByeRequest) (*farewellpd.GoodByeResponse, error) {
	msg := "Goodbye, " + req.Name + "!"
	log.Printf("BigGoodBye() called: %s", msg)

	return &farewellpd.GoodByeResponse{
		Message: msg,
	}, nil
}

func main() {
	// TLS certificate files
	// These must be present in the current directory
	certFile := "cert.pem"
	keyFile := "key.pem"

	// Server port
	port := "50051"

	// ============================================================
	// Step 1: Create TCP Listener
	// ============================================================
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	// ============================================================
	// Step 2: Load TLS Credentials
	// ============================================================
	// Load the server's certificate and private key
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	// ============================================================
	// Step 3: Create gRPC Server with TLS
	// ============================================================
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	// ============================================================
	// Step 4: Register Services
	// ============================================================
	// A single server instance can implement multiple services
	svc := &server{}
	pb.RegisterCalculatorServer(grpcServer, svc)
	pb.RegisterGreeterServer(grpcServer, svc)
	farewellpd.RegisterAufWiedersehenServer(grpcServer, svc)

	log.Printf("Secure gRPC Server listening on port %s", port)
	log.Println("Services registered: Calculator, Greeter, AufWiedersehen")

	// ============================================================
	// Step 5: Start Serving
	// ============================================================
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
