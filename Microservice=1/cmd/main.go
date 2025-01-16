package main

import (
	"fmt"
	"log"
	"methodOne/pkg/config"
	"methodOne/pkg/di"
	"methodOne/pkg/pb"
	"net"

	"google.golang.org/grpc"
)

func main() {
	
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	server, err := di.InitializeAuthService(cfg)
	if err != nil {
		log.Fatalf("Error initializing auth service: %v", err)
	}

	port := fmt.Sprintf(":%s", cfg.PortMngr.PortNo)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start TCP listener: %v", err)
	}

	log.Printf("Auth Service started on port: %s", cfg.PortMngr.PortNo)

	
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
