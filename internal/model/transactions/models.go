package transactions

import "time"

type (
	CheckoutModel struct {
		ID            int64     `db:"id"`
		TransactionID int64     `db:"transaction_id"`
		ProductSKU    []string  `db:"product_sku"`
		CreatedAt     time.Time `db:"created_at"`
		UpdatedAt     time.Time `db:"updated_at"`
	}

	TransactionModel struct {
		ID                  int64     `db:"id"`
		UserID              int64     `db:"user_id"`
		Subtotal            float64   `db:"subtotal"`
		Discount            float64   `db:"discount"`
		DiscountExplanation string    `db:"discount_explanation"`
		Total               float64   `db:"total"`
		CreatedAt           time.Time `db:"created_at"`
		UpdatedAt           time.Time `db:"updated_at"`
	}

	ProductModel struct {
		ID           int64     `db:"id"`
		SKU          string    `db:"sku"`
		Name         string    `db:"name"`
		Price        float64   `db:"price"`
		InventoryQty int64     `db:"inventory_qty"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
	}
)

type (
	CheckoutRequest struct {
		UserID     int64    `json:"user_id"`
		ProductSKU []string `json:"product_sku"`
	}

	Product struct {
		SKU   string  `json:"sku"`
		Price float64 `json:"price"`
	}

	ProductResponse struct {
		Product []Product `json:"products"`
	}
)
