package main

import (
	"context"
	"log"
	mainapipb "simplegrpcserver/proto/gen" // Derlenen proto paketleri
	farewellpb "simplegrpcserver/proto/gen/farewell"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	// --- TLS Sertifikasını yükle ---
	// NewClientTLSFromFile: client tarafında TLS doğrulaması için sertifika verir
	// İkinci parametre serverNameOverride -> TLS handshake sırasında CN eşleşmesi için "localhost"
	creds, err := credentials.NewClientTLSFromFile("cert.pem", "localhost")
	if err != nil {
		log.Fatalln("Failed to load TLS certificate:", err)
	}

	// --- gRPC Sunucusuna TLS ile bağlan ---
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer conn.Close()

	// --- RPC client örneklerini oluştur ---
	calculatorClient := mainapipb.NewCalculatorClient(conn)
	greeterClient := mainapipb.NewGreeterClient(conn)
	farewellClient := farewellpb.NewAufWiedersehenClient(conn)

	// --- Zaman aşımı ekleyerek güvenli bir context oluştur ---
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// --- Calculator Servisine istek ---
	addReq := &mainapipb.AddRequest{
		A: 10,
		B: 20,
	}

	addRes, err := calculatorClient.Add(ctx, addReq)
	if err != nil {
		log.Fatalln("Add RPC failed:", err)
	}

	// --- Farewell Servisine istek ---
	reqGoodBye := &farewellpb.GoodByeRequest{
		Name: "Yasin",
	}
	goodbyeRes, err := farewellClient.BigGoodBye(ctx, reqGoodBye)
	if err != nil {
		log.Fatalln("Goodbye RPC failed:", err)
	}

	// --- Greeter Servisine istek ---
	greetReq := &mainapipb.HelloRequest{
		Name: "Yasin",
	}

	greetRes, err := greeterClient.Greet(ctx, greetReq)
	if err != nil {
		log.Fatalln("Could not greet:", err)
	}

	// --- RPC çıktıları ---
	log.Println("Sum:", addRes.Sum)
	log.Println("Greeting message:", greetRes.Message)
	log.Println("Goodbye message:", goodbyeRes.Message)

	// --- Bağlantı durumunu yazdır ---
	state := conn.GetState()
	log.Println("Connection State:", state)
}
