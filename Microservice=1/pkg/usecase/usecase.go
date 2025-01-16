package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"methodOne/pkg/db"
	"methodOne/pkg/model"
	interfaces "methodOne/pkg/repo/interface"
	interface_usecase "methodOne/pkg/usecase/interfaces"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserUsecase struct {
	Repo  interfaces.IUserRepo
	Redis *redis.Client
	Ctx   context.Context
}

func NewUserUsecase(repo interfaces.IUserRepo, redis *redis.Client, ctx context.Context) interface_usecase.IUserUsecase {
	return &UserUsecase{Repo: repo, Redis: redis, Ctx: ctx}
}

func (u *UserUsecase) CreateUser(user model.User) error {
	if err := u.Repo.CreateUser(user); err != nil {
		return err
	}
	// Invalidate cache for all users
	db.DeleteFeedEntry(u.Ctx, u.Redis, "GetAllUser")
	return nil
}

func (u *UserUsecase) GetUserByID(id uint64) (*model.User, error) {
	cacheKey := fmt.Sprintf("GetUser:%d", id) // Key specific to user ID

	// Attempt to retrieve cached data
	cachedUserJSON, err := u.Redis.Get(u.Ctx, cacheKey).Result()
	if err == redis.Nil {
		// Cache miss - fetch from DB
		user, err := u.Repo.GetUserByID(id)
		if err != nil {
			return nil, err
		}

		// Cache the fetched data (serialize to JSON)
		userJSON, err := json.Marshal(user)
		if err != nil {
			return nil, fmt.Errorf("error serializing user data: %w", err)
		}
		err = u.Redis.Set(u.Ctx, cacheKey, userJSON, 10*time.Minute).Err()
		if err != nil {
			fmt.Printf("Failed to cache user data: %v\n", err)
		}

		return user, nil
	} else if err != nil {
		// Redis error (not a cache miss)
		return nil, fmt.Errorf("redis error: %w", err)
	}

	// Cache hit - deserialize the data
	var user model.User
	err = json.Unmarshal([]byte(cachedUserJSON), &user)
	if err != nil {
		return nil, fmt.Errorf("error deserializing user data: %w", err)
	}

	return &user, nil
}

func (u *UserUsecase) UpdateUser(user *model.User) error {
	if err := u.Repo.UpdateUser(user); err != nil {
		return err
	}

	// Invalidate cache for specific user and all users
	cacheKeyByID := fmt.Sprintf("GetUser:%d", user.ID)
	db.DeleteFeedEntry(u.Ctx, u.Redis, cacheKeyByID)
	db.DeleteFeedEntry(u.Ctx, u.Redis, "GetAllUser")

	// Optionally, you could cache the updated user here if necessary
	return nil
}

func (u *UserUsecase) DeleteUser(id uint64) error {
	if err := u.Repo.DeleteUser(id); err != nil {
		return err
	}

	// Invalidate cache for specific user and all users
	cacheKeyByID := fmt.Sprintf("GetUser:%d", id)
	db.DeleteFeedEntry(u.Ctx, u.Redis, cacheKeyByID)
	db.DeleteFeedEntry(u.Ctx, u.Redis, "GetAllUser")
	return nil
}

func (u *UserUsecase) ListAllUsers() ([]model.User, error) {
	cacheKey := "GetAllUser"

	// Attempt to retrieve cached data
	cachedUsersJSON, err := u.Redis.Get(u.Ctx, cacheKey).Result()
	if err == redis.Nil {
		// Cache miss - fetch from DB
		users, err := u.Repo.ListAllUsers()
		if err != nil {
			return nil, err
		}

		// Cache the fetched data (serialize to JSON)
		usersJSON, err := json.Marshal(users)
		if err != nil {
			return nil, fmt.Errorf("error serializing users data: %w", err)
		}
		err = u.Redis.Set(u.Ctx, cacheKey, usersJSON, 10*time.Minute).Err()
		if err != nil {
			fmt.Printf("Failed to cache users data: %v\n", err)
		}

		return users, nil
	} else if err != nil {
		// Redis error (not a cache miss)
		return nil, fmt.Errorf("redis error: %w", err)
	}

	// Cache hit - deserialize the data
	var users []model.User
	err = json.Unmarshal([]byte(cachedUsersJSON), &users)
	if err != nil {
		return nil, fmt.Errorf("error deserializing users data: %w", err)
	}

	return users, nil
}
