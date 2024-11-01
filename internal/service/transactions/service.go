package transactions

import (
	"context"

	"github.com/tegarpratama/checkout-service/internal/configs"
	"github.com/tegarpratama/checkout-service/internal/model/transactions"
	"github.com/tegarpratama/checkout-service/internal/model/users"
)

type transactionRepo interface {
	GetProductsBySKU(ctx context.Context, productSku []string) (transactions.ProductResponse, error)
	GetUserByID(ctx context.Context, userID int64) (*users.UserModel, error)
	StoreTransaction(ctx context.Context, model transactions.TransactionModel) (int64, error)
	StoreCheckout(ctx context.Context, model transactions.CheckoutModel) error
}

type service struct {
	cfg             *configs.Config
	transactionRepo transactionRepo
}

func NewService(cfg *configs.Config, transactionRepo transactionRepo) *service {
	return &service{
		cfg:             cfg,
		transactionRepo: transactionRepo,
	}
}
