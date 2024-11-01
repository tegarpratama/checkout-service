package users

import (
	"context"
	"errors"
	"time"

	"github.com/tegarpratama/checkout-service/internal/model/users"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) CreateUser(ctx context.Context, req users.CreateUserRequest) (int64, error) {
	user, err := s.userRepo.GetUser(ctx, req.Email)
	if err != nil {
		return 0, err
	}

	if user != nil {
		return 0, errors.New("email already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	now := time.Now()
	model := users.UserModel{
		Email:     req.Email,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
	}

	userID, err := s.userRepo.CreateUser(ctx, model)
	if err != nil {
		return 0, nil
	}

	return userID, nil
}
