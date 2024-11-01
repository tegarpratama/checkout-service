package users

import (
	"context"

	"github.com/tegarpratama/checkout-service/internal/configs"
	"github.com/tegarpratama/checkout-service/internal/model/users"
)

type userRepository interface {
	CreateUser(ctx context.Context, model users.UserModel) (int64, error)
	GetUser(ctx context.Context, email string) (*users.UserModel, error)
}

type service struct {
	cfg      *configs.Config
	userRepo userRepository
}

func NewService(cfg *configs.Config, userRepo userRepository) *service {
	return &service{
		cfg:      cfg,
		userRepo: userRepo,
	}
}
