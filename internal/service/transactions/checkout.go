package transactions

import (
	"context"
	"errors"
	"math"
	"strings"
	"time"

	"github.com/tegarpratama/checkout-service/internal/model/transactions"
)

const (
	SKUMacbookPro = "43N23P"
	SKURaspberry  = "234234"
	SKUGoogleHome = "120P90"
	SKUAlexa      = "A304SD"
)

type productDetail struct {
	Qty   int16
	Price float64
}

type calculatePrice struct {
	TotalDiscount       float64
	DiscountInformation []string
	TotalAfterDiscount  float64
}

func (s *service) Checkout(ctx context.Context, req transactions.CheckoutRequest) (float64, error) {
	now := time.Now()
	subtotal := 0.00

	// Check user exists
	user, err := s.transactionRepo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return 0.00, err
	}

	if user == nil {
		return 0.00, errors.New("user not found")
	}

	// Get products information
	productMaster, err := s.transactionRepo.GetProductsBySKU(ctx, req.ProductSKU)
	if err != nil {
		return 0.00, err
	}

	// Mapping product
	productMap := make(map[string]transactions.Product)
	for _, product := range productMaster.Product {
		productMap[product.SKU] = product
	}

	productsBoughtDiscount := map[string]*productDetail{
		SKUMacbookPro: {Qty: 0, Price: productMap[SKUMacbookPro].Price},
		SKURaspberry:  {Qty: 0, Price: productMap[SKURaspberry].Price},
		SKUGoogleHome: {Qty: 0, Price: productMap[SKUGoogleHome].Price},
		SKUAlexa:      {Qty: 0, Price: productMap[SKUAlexa].Price},
	}

	for _, checkoutSKU := range req.ProductSKU {
		if product, exists := productMap[checkoutSKU]; exists {
			productsBoughtDiscount[checkoutSKU].Qty++
			subtotal += product.Price
		}
	}

	// Apply discount if exists
	result := applyDiscounts(subtotal, productsBoughtDiscount)

	// Rounded tototal
	total := math.Round(result.TotalAfterDiscount*100) / 100

	// Store data transaction
	valueInsertTransaction := transactions.TransactionModel{
		UserID:              req.UserID,
		Subtotal:            subtotal,
		Discount:            result.TotalDiscount,
		DiscountExplanation: strings.Join(result.DiscountInformation, ", "),
		Total:               total,
		CreatedAt:           now,
		UpdatedAt:           now,
	}

	transactionID, err := s.transactionRepo.StoreTransaction(ctx, valueInsertTransaction)
	if err != nil {
		return 0.00, err
	}

	// Store data checkout
	valueInsertCheckout := transactions.CheckoutModel{
		TransactionID: transactionID,
		ProductSKU:    req.ProductSKU,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	err = s.transactionRepo.StoreCheckout(ctx, valueInsertCheckout)
	if err != nil {
		return 0.00, err
	}

	return total, nil
}

func applyDiscounts(total float64, productsBought map[string]*productDetail) *calculatePrice {
	discountTotal := 0.00
	var discountExplanation []string

	// Rule: user buy MacbookPro with Raspberry
	macbookPro, macExists := productsBought[SKUMacbookPro]
	raspberry, rasExists := productsBought[SKURaspberry]

	if macExists && macbookPro != nil &&
		rasExists && raspberry != nil &&
		productsBought[SKUMacbookPro].Qty > 0 &&
		productsBought[SKURaspberry].Qty > 0 &&
		productsBought[SKUMacbookPro].Qty == productsBought[SKURaspberry].Qty {
		discount := float64(productsBought[SKURaspberry].Qty) * productsBought[SKURaspberry].Price
		discountExplanation = append(discountExplanation, "Each sale of a MacBook Pro comes with a free Raspberry Pi B")

		total -= discount
		discountTotal += discount
	}

	// Rule: buy 3 Google Home devices
	googleHome, googleHomeExists := productsBought[SKUGoogleHome]

	if googleHomeExists && googleHome != nil &&
		productsBought[SKUGoogleHome].Qty >= 3 {
		discount := productsBought[SKUGoogleHome].Price
		discountExplanation = append(discountExplanation, "Buy 3 Google Homes for the price of 2")

		total -= discount
		discountTotal += discount
	}

	// Rule: buy 3 or more Alexa speakers
	alexa, alexaExists := productsBought[SKUAlexa]

	if alexaExists && alexa != nil &&
		productsBought[SKUAlexa].Qty >= 3 {
		discount := (productsBought[SKUAlexa].Price * 0.1) * float64(productsBought[SKUAlexa].Qty)
		discountExplanation = append(discountExplanation, "Buying more than 3 Alexa Speakers will get a 10% discount on all Alexa speakers")

		total -= discount
		discountTotal += discount
	}

	return &calculatePrice{
		TotalDiscount:       discountTotal,
		DiscountInformation: discountExplanation,
		TotalAfterDiscount:  total,
	}
}
