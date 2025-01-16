package interfaces

import "methodOne/pkg/model"

type IUserRepo interface {
	CreateUser(user model.User) error
	GetUserByID(id uint64) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint64) error
	ListAllUsers() ([]model.User, error)
}
