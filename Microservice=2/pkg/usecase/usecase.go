package usecase

import (
	"api-gateway/pkg/Microservice/client"
	queue "api-gateway/pkg/Queue"
	"context"
	"time"
)

type UserService struct {
	userClient *client.UserClient
	taskQueue  *queue.TaskQueue
}

// NewUserService creates a new UserService instance.
func NewUserService(userClient *client.UserClient, taskQueue *queue.TaskQueue) *UserService {
	return &UserService{
		userClient: userClient,
		taskQueue:  taskQueue,
	}
}

// Method1: Sequential execution via task queue.
func (s *UserService) Method1(ctx context.Context, waitTime int) ([]string, error) {
	resultChan := make(chan []string, 1)
	errChan := make(chan error, 1)

	// Add the task to the task queue for sequential processing.
	s.taskQueue.AddTask(func() {
		defer close(resultChan)
		defer close(errChan)

		// Fetch users using the client.
		users, err := s.userClient.ListUsers(ctx)
		if err != nil {
			errChan <- err
			return
		}

		// Collect user names.
		userNames := make([]string, len(users))
		for i, user := range users {
			userNames[i] = user.Name
		}

		// Simulate work by sleeping.
		time.Sleep(time.Duration(waitTime) * time.Second)

		// Send result to the result channel.
		resultChan <- userNames
	})

	// Retrieve results from the channels.
	select {
	case res := <-resultChan:
		return res, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		// Handle context timeout or cancellation.
		return nil, ctx.Err()
	}
}

// Method2: Parallel execution without queue involvement.
func (s *UserService) Method2(ctx context.Context, waitTime int) ([]string, error) {
	// Fetch users using the client.
	users, err := s.userClient.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	// Collect user names.
	userNames := make([]string, len(users))
	for i, user := range users {
		userNames[i] = user.Name
	}

	// Simulate work by sleeping.
	time.Sleep(time.Duration(waitTime) * time.Second)

	return userNames, nil
}
