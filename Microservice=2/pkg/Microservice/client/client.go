package client

import (
	"api-gateway/pkg/config"
	"api-gateway/pkg/pb"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type UserClient struct {
	pb.UserServiceClient
}

// InitUserClient initializes a gRPC client for the UserService.
func InitUserClient(cfg *config.Config) (*UserClient, error) {
	conn, err := grpc.Dial(cfg.UserSvcUrl, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %w", err)
	}

	return &UserClient{
		UserServiceClient: pb.NewUserServiceClient(conn),
	}, nil
}

// ListUsers calls the ListAllUsers RPC and returns a list of users.
func (uc *UserClient) ListUsers(ctx context.Context) ([]*pb.User, error) {
	// Call the gRPC ListAllUsers method
	resp, err := uc.ListAllUsers(ctx, &pb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	// Return the list of users from the response
	return resp.Users, nil
}
