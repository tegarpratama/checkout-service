package products

import "time"

type (
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
