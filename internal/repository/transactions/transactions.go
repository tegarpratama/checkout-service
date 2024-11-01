package transactions

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/tegarpratama/checkout-service/internal/model/transactions"
	"github.com/tegarpratama/checkout-service/internal/model/users"
)

func (r *repository) GetProductsBySKU(ctx context.Context, productSku []string) (transactions.ProductResponse, error) {
	var response transactions.ProductResponse

	placeholders := make([]string, len(productSku))
	args := make([]interface{}, len(productSku))
	for i, sku := range productSku {
		placeholders[i] = "?"
		args[i] = sku
	}

	query := fmt.Sprintf(`SELECT sku, price FROM products WHERE sku IN (%s)`, strings.Join(placeholders, ", "))
	rows, err := r.db.QueryContext(ctx, query, args...)

	if err != nil {
		return response, nil
	}

	defer rows.Close()
	data := make([]transactions.Product, 0)

	for rows.Next() {
		var model transactions.Product

		err := rows.Scan(&model.SKU, &model.Price)
		if err != nil {
			return response, nil
		}

		data = append(data, transactions.Product{
			SKU:   model.SKU,
			Price: model.Price,
		})
	}

	response.Product = data

	return response, nil
}

func (r *repository) GetUserByID(ctx context.Context, userID int64) (*users.UserModel, error) {
	var response users.UserModel

	query := `SELECT id FROM users WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, userID)
	err := row.Scan(&response.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &response, nil
}

func (r *repository) StoreTransaction(ctx context.Context, model transactions.TransactionModel) (int64, error) {
	query := `INSERT INTO transactions (user_id, subtotal, discount, discount_explanation, total, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := r.db.ExecContext(ctx, query, model.UserID, model.Subtotal, model.Discount, model.DiscountExplanation, model.Total, model.CreatedAt, model.UpdatedAt)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) StoreCheckout(ctx context.Context, model transactions.CheckoutModel) error {
	query := `INSERT INTO checkouts (transaction_id, product_sku, created_at, updated_at) VALUES `
	values := make([]string, len(model.ProductSKU))
	args := make([]interface{}, 0, len(model.ProductSKU)*4)

	for i, sku := range model.ProductSKU {
		values[i] = "(?, ?, ?, ?)"
		args = append(args, model.TransactionID, sku, model.CreatedAt, model.UpdatedAt)
	}

	query += strings.Join(values, ", ")

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
