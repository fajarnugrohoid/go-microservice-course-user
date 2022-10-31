package service

import (
	"context"
	"errors"
	"user-service/internal/app/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, user *domain.User) error
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo}
}

func (us UserService) Login(ctx context.Context, param domain.LoginParam) (string, error) {
	user, err := us.userRepo.GetByEmail(ctx, param.Email)
	if err != nil {
		return "", errors.New("wrong email/password")
	}

	if err := user.ComparePassword(param.Password); err != nil {
		return "", errors.New("wrong email/password")
	}
	return generateJWT(user.ID)
}

func (us UserService) Register(ctx context.Context, param domain.RegisterParam) (string, error) {
	user := domain.NewUser(param.Name, param.Email, param.Password)
	if err := us.userRepo.Create(ctx, user); err != nil {
		return "", errors.New("failed create user")
	}
	return generateJWT(user.ID)
}
