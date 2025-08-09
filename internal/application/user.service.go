package application

import (
	"context"
	"github.com/Alao-Abiodun/lender-api/internal/domain/user"
)

type UserService struct {
	userRepo UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *user.User) error
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{ userRepo }
}

func (userService *UserService) RegisterUser(ctx context.Context, user *user.User) error {
	return userService.userRepo.CreateUser(ctx, user)
}