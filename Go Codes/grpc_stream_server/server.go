package main

import (
	mainpb "grpcstream/proto/gen"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

// ---Calculator Servisinin impelamtasyonu
type server struct {
	mainpb.UnimplementedCalculatorServer
}

// ---Server-side streaming---
func (*server) GenerateFibonacci(req *mainpb.FibonacciRequest, stream mainpb.Calculator_GenerateFibonacciServer) error {
	n := req.N
	a, b := 0, 1

	//---N tane itesaryon şekilnde clien'ta değer gönderilir
	for i := 0; i < int(n); i++ {
		err := stream.Send(&mainpb.FibonacciResponse{ //---Clienta anlın mesaj yollarım---
			Number: int32(a),
		})
		if err != nil {
			return err
		}
		log.Println("Sent Number:",a)
		a, b = b, a+b
		time.Sleep(time.Second) //---1 saniyelik aralıklarla---
	}
	return nil
}

func main() {
	//---Clientın bağlanacağı port---
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalln(err)
	}

	//---Bos bir gRPC instance oluşturulur---
	grpcServer := grpc.NewServer()

	//---Calcualtor servisini gRPC server'a repister etme---
	mainpb.RegisterCalculatorServer(grpcServer, &server{})

	//---Sunucuyu çalıştırma---
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
