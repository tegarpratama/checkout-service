package seeder

import (
	"database/sql"
	"log"

	"github.com/tegarpratama/checkout-service/internal/model/products"
)

func SeedProducts(db *sql.DB) {
	productInserted := []products.ProductModel{
		{
			SKU:          "120P90",
			Name:         "Google Home",
			Price:        49.99,
			InventoryQty: 10,
		},
		{
			SKU:          "43N23P",
			Name:         "MacBook Pro",
			Price:        5399.99,
			InventoryQty: 5,
		},
		{
			SKU:          "A304SD",
			Name:         "Alexa Speaker",
			Price:        109.50,
			InventoryQty: 10,
		},
		{
			SKU:          "234234",
			Name:         "Raspberry Pi B",
			Price:        30.00,
			InventoryQty: 2,
		},
	}

	for _, product := range productInserted {
		var productModel products.ProductModel
		rows := db.QueryRow(`SELECT sku FROM products WHERE sku = ?`, product.SKU)
		err := rows.Scan(&productModel.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				_, err := db.Exec(`INSERT INTO products (sku, name, price, inventory_qty) VALUES (?, ?, ?, ?)`, product.SKU, product.Name, product.Price, product.InventoryQty)
				if err != nil {
					log.Fatalf("could not insert product %s: %v", product.Name, err)
				}
			}
		}
	}

	log.Println("finish seed products data")
}
