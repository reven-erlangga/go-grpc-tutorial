package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"go_grpc/cmd/config"
	"go_grpc/cmd/services"
	pb "go_grpc/pb"
)

const (
	port = ":50051"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listend %v", err.Error())
	}

	db := config.ConnectDatabase()

	grpcServer := grpc.NewServer()
	productService := services.ProductService{DB: db}
	pb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}
