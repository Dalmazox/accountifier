package main

import (
	"log"
	"net"

	authv1 "github.com/dalmazox/accountifier/gen/go/auth/v1"
	grpcservices "github.com/dalmazox/accountifier/internal/grpc_services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao iniciar o listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	authv1.RegisterAuthServiceServer(grpcServer, grpcservices.NewAuthServiceServer())

	reflection.Register(grpcServer)

	log.Printf("Servidor gRPC rodando na porta :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Falha ao iniciar o servidor gRPC: %v", err)
	}
}
