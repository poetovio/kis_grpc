package main

import (
	"grpcserver/dopust"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Napaka pri poslusanju na portu 50051: %v", err)
	}

	server := dopust.Server{}

	grpcServer := grpc.NewServer()

	dopust.RegisterDopustServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}

}
