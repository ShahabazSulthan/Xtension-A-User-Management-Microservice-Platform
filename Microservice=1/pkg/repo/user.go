package repo

import (
	"fmt"
	"methodOne/pkg/model"
	interfaces "methodOne/pkg/repo/interface"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.IUserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) CreateUser(user model.User) error {
	if err := u.DB.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (u *UserRepo) GetUserByID(id uint64) (*model.User, error) {
	var user model.User
	if err := u.DB.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("failed to find user with ID %d: %w", id, err)
	}
	return &user, nil
}

func (u *UserRepo) UpdateUser(user *model.User) error {
	if err := u.DB.Model(&model.User{}).
		Where("email = ?", user.Email).
		Updates(map[string]interface{}{
			"name":       user.Name,
			"phone":      user.Phone,
			"updated_at": gorm.Expr("NOW()"),
		}).Error; err != nil {
		return fmt.Errorf("failed to update user with email %s: %w", user.Email, err)
	}
	return nil
}

func (u *UserRepo) DeleteUser(id uint64) error {
	if err := u.DB.Delete(&model.User{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user with ID %d: %w", id, err)
	}
	return nil
}

func (u *UserRepo) ListAllUsers() ([]model.User, error) {
	var users []model.User
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch all users: %w", err)
	}
	return users, nil
}
