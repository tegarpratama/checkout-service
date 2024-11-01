package transactions

import (
	"context"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tegarpratama/checkout-service/internal/configs"
	"github.com/tegarpratama/checkout-service/internal/model/transactions"
	"github.com/tegarpratama/checkout-service/internal/model/users"
)

type mockTransactionRepo struct {
	mock.Mock
}

func (m *mockTransactionRepo) GetProductsBySKU(ctx context.Context, productSku []string) (transactions.ProductResponse, error) {
	args := m.Called(ctx, productSku)
	return args.Get(0).(transactions.ProductResponse), args.Error(1)
}

func (m *mockTransactionRepo) GetUserByID(ctx context.Context, userID int64) (*users.UserModel, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*users.UserModel), args.Error(1)
}

func (m *mockTransactionRepo) StoreTransaction(ctx context.Context, model transactions.TransactionModel) (int64, error) {
	args := m.Called(ctx, model)
	return args.Get(0).(int64), args.Error(1)
}

func (m *mockTransactionRepo) StoreCheckout(ctx context.Context, model transactions.CheckoutModel) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func TestCheckout_Success(t *testing.T) {
	mockRepo := new(mockTransactionRepo)
	cfg := &configs.Config{}
	svc := NewService(cfg, mockRepo)

	mockRepo.On("GetUserByID", mock.Anything, int64(1)).Return(&users.UserModel{ID: 1}, nil)

	mockRepo.On("GetProductsBySKU", mock.Anything, mock.Anything).Return(transactions.ProductResponse{
		Product: []transactions.Product{
			{SKU: SKUMacbookPro, Price: 1000.00},
			{SKU: SKURaspberry, Price: 50.00},
		},
	}, nil)

	mockRepo.On("StoreTransaction", mock.Anything, mock.Anything).Return(int64(1), nil)

	mockRepo.On("StoreCheckout", mock.Anything, mock.Anything).Return(nil)

	req := transactions.CheckoutRequest{
		UserID:     1,
		ProductSKU: []string{SKUMacbookPro, SKURaspberry},
	}

	total, err := svc.Checkout(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, 1000.00, total)
	mockRepo.AssertExpectations(t)
}

func TestDiscountBuyMacbookWithRaspberry(t *testing.T) {
	productsBought := map[string]*productDetail{
		SKUMacbookPro: {Qty: 1, Price: 5399.99},
		SKURaspberry:  {Qty: 1, Price: 30.00},
	}

	total := productsBought[SKUMacbookPro].Price + productsBought[SKURaspberry].Price
	result := applyDiscounts(total, productsBought)

	expectedTotal := productsBought[SKUMacbookPro].Price
	assert.Equal(t, expectedTotal, result.TotalAfterDiscount)
}

func TestDiscountBuyThreeGoogleHome(t *testing.T) {
	productsBought := map[string]*productDetail{
		SKUGoogleHome: {Qty: 3, Price: 49.99},
	}

	total := productsBought[SKUGoogleHome].Price * float64(productsBought[SKUGoogleHome].Qty)
	result := applyDiscounts(total, productsBought)

	expectedTotal := productsBought[SKUGoogleHome].Price * float64(productsBought[SKUGoogleHome].Qty-1)
	totalRounded := math.Round(result.TotalAfterDiscount*100) / 100

	assert.Equal(t, expectedTotal, totalRounded)
}

func TestDiscountBuyThreeAlexa(t *testing.T) {
	productsBought := map[string]*productDetail{
		SKUAlexa: {Qty: 3, Price: 109.50},
	}

	total := productsBought[SKUAlexa].Price * float64(productsBought[SKUAlexa].Qty)
	result := applyDiscounts(total, productsBought)

	expectedTotal := total - (productsBought[SKUAlexa].Price*0.1)*3
	totalRounded := math.Round(result.TotalAfterDiscount*100) / 100

	assert.Equal(t, expectedTotal, totalRounded)
}
