package users

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tegarpratama/checkout-service/internal/configs"
	"github.com/tegarpratama/checkout-service/internal/model/users"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) CreateUser(ctx context.Context, model users.UserModel) (int64, error) {
	args := m.Called(ctx, model)
	return args.Get(0).(int64), args.Error(1)
}

func (m *mockUserRepository) GetUser(ctx context.Context, email string) (*users.UserModel, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*users.UserModel), args.Error(1)
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mockUserRepository)
	cfg := &configs.Config{}
	svc := NewService(cfg, mockRepo)

	req := users.CreateUserRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	mockRepo.On("GetUser", mock.Anything, req.Email).Return((*users.UserModel)(nil), nil)
	mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("users.UserModel")).Return(int64(1), nil)

	userID, err := svc.CreateUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), userID)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_EmailAlreadyExists(t *testing.T) {
	mockRepo := new(mockUserRepository)
	cfg := &configs.Config{}
	svc := NewService(cfg, mockRepo)

	req := users.CreateUserRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	existingUser := &users.UserModel{Email: req.Email}
	mockRepo.On("GetUser", mock.Anything, req.Email).Return(existingUser, nil)

	userID, err := svc.CreateUser(context.Background(), req)

	assert.Error(t, err)
	assert.Equal(t, "email already exists", err.Error())
	assert.Equal(t, int64(0), userID)
	mockRepo.AssertExpectations(t)
}
