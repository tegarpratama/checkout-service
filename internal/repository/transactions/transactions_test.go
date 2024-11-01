package transactions

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tegarpratama/checkout-service/internal/model/transactions"
)

func TestGetProductsBySKU(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	repo := NewRepository(db)

	productSkus := []string{"sku1", "sku2"}
	rows := sqlmock.NewRows([]string{"sku", "price"}).
		AddRow("sku1", 100).
		AddRow("sku2", 200)

	mock.ExpectQuery(`SELECT sku, price FROM products WHERE sku IN \(\?, \?\)`).
		WithArgs("sku1", "sku2").
		WillReturnRows(rows)

	response, err := repo.GetProductsBySKU(context.Background(), productSkus)

	assert.NoError(t, err)
	assert.Len(t, response.Product, 2)
	assert.Equal(t, "sku1", response.Product[0].SKU)
	assert.Equal(t, float64(100), response.Product[0].Price)
	assert.Equal(t, "sku2", response.Product[1].SKU)
	assert.Equal(t, float64(200), response.Product[1].Price)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	repo := NewRepository(db)

	userID := int64(1)
	mock.ExpectQuery(`SELECT id FROM users WHERE id = \?`).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))

	user, err := repo.GetUserByID(context.Background(), userID)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.ID)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestStoreTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	repo := NewRepository(db)

	model := transactions.TransactionModel{
		UserID:              1,
		Subtotal:            300,
		Discount:            50,
		DiscountExplanation: "Promo",
		Total:               250,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	mock.ExpectExec(`INSERT INTO transactions`).
		WithArgs(model.UserID, model.Subtotal, model.Discount, model.DiscountExplanation, model.Total, model.CreatedAt, model.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := repo.StoreTransaction(context.Background(), model)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
