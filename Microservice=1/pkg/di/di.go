package di

import (
	"context"
	"fmt"
	"methodOne/pkg/config"
	"methodOne/pkg/db"
	"methodOne/pkg/repo"
	"methodOne/pkg/server"
	"methodOne/pkg/usecase"
)

func InitializeAuthService(config *config.Config) (*server.UserService, error) {
	DB, err := db.ConnectDatabase(&config.DB)
	if err != nil {
		fmt.Println("Error in connecting database from Dependency Injection")
		return nil, err
	}

	redisClient, err := db.NewRedisHelper(config.Redis)
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return nil, err
	}

	ctx := context.Background()

	repository := repo.NewRepository(DB)
	usecase := usecase.NewUserUsecase(repository, redisClient.Client, ctx)
	service := server.NewUserService(usecase)

	return service, nil
}
