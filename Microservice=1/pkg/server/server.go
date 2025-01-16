package server

import (
	"context"
	"methodOne/pkg/model"
	"methodOne/pkg/pb"
	interface_usecase "methodOne/pkg/usecase/interfaces"
)

type UserService struct {
	Usecase interface_usecase.IUserUsecase
	pb.UnimplementedUserServiceServer
}

// NewUserService initializes a new UserService with the given usecase.
func NewUserService(usecase interface_usecase.IUserUsecase) *UserService {
	return &UserService{
		Usecase: usecase,
	}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}
	if err := s.Usecase.CreateUser(*user); err != nil {
		return &pb.CreateUserResponse{
			Id:      0,
			Message: "Failed to create user: " + err.Error(),
		}, err
	}
	return &pb.CreateUserResponse{
		Id:      user.ID, // Assuming the `CreateUser` method sets the user's ID.
		Message: "User created successfully",
	}, nil
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := s.Usecase.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByIDResponse{
		User: &pb.User{
			Id:    uint64(user.ID),
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		},
	}, nil
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := &model.User{
		ID:    req.Id,
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}
	if err := s.Usecase.UpdateUser(user); err != nil {
		return &pb.UpdateUserResponse{
			Message: "Failed to update user: " + err.Error(),
		}, err
	}
	return &pb.UpdateUserResponse{
		Message: "User updated successfully",
	}, nil
}

// DeleteUser deletes a user by their ID.
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if err := s.Usecase.DeleteUser(req.Id); err != nil {
		return &pb.DeleteUserResponse{
			Message: "Failed to delete user: " + err.Error(),
		}, err
	}
	return &pb.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}

// ListAllUsers retrieves all users in the system.
func (s *UserService) ListAllUsers(ctx context.Context, req *pb.Empty) (*pb.ListAllUsersResponse, error) {
	users, err := s.Usecase.ListAllUsers()
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:    uint64(user.ID),
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		})
	}

	return &pb.ListAllUsersResponse{
		Users: pbUsers,
	}, nil
}
