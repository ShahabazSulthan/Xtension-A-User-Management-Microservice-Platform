package main

import (
	"api-gateway/pkg/Microservice/client"
	"api-gateway/pkg/Microservice/handler"
	"api-gateway/pkg/Microservice/routes"
	queue "api-gateway/pkg/Queue"
	"api-gateway/pkg/config"
	"api-gateway/pkg/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	
	userClient, err := client.InitUserClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client for UserService: %v", err)
	}

	
	taskQueue := queue.NewTaskQueue(100) 
	
	userService := usecase.NewUserService(userClient, taskQueue)

	
	userHandler := handler.NewUserHandler(userClient, userService)

	
	router := gin.Default()
	router.Use(gin.Recovery())

	
	routes.UserRoutes(router, userHandler)

	
	if cfg.Port == "" {
		log.Fatalf("Configuration error: Port must be specified")
	}

	
	serverAddr := "0.0.0.0" + cfg.Port 
	log.Printf("Server is running on %s", serverAddr)

	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
