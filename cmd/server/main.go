package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dalmazox/accountifier/config"
	authv1 "github.com/dalmazox/accountifier/gen/go/auth/v1"
	grpcservices "github.com/dalmazox/accountifier/internal/grpc_services"
	"github.com/dalmazox/accountifier/internal/repositories"
	"github.com/dalmazox/accountifier/internal/usecases"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("could not load config file: %v", err)
	}

	db, err := sqlx.Connect("postgres", cfg.Database.ConnectionString)
	if err != nil {
		log.Fatalf("fail to connect database")
	}
	defer db.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatalf("fail to listen to port: %v", err)
	}

	grpcServer := grpc.NewServer()

	userRepo := repositories.NewUserRepository(db)
	userTokenRepo := repositories.NewUserTokenRepository(db)
	loginUseCase := usecases.NewLoginUseCase(cfg, userRepo, userTokenRepo)
	authv1.RegisterAuthServiceServer(grpcServer, grpcservices.NewAuthServiceServer(loginUseCase))

	reflection.Register(grpcServer)

	log.Printf("running server in port :%d", cfg.Server.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start server in port :%v", err)
	}
}
