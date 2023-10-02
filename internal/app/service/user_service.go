package service

import (
	"context"

	"github.com/cetalogy/internal/app/model"
	"github.com/cetalogy/internal/app/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (us *UserService) CreateUser(ctx context.Context, user *model.User) error {
	return us.userRepository.CreateUser(ctx, user)
}

func (us *UserService) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	return us.userRepository.GetUserByID(ctx, userID)
}

func (us *UserService) GetUsers(ctx context.Context) (*[]model.User, error) {
	return us.userRepository.GetUsers(ctx)
}

func (us *UserService) DeleteUserByID(ctx context.Context, userID string) error {
	return us.userRepository.DeleteUserByID(ctx, userID)
}
