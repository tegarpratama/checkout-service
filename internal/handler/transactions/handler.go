package transactions

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tegarpratama/checkout-service/internal/model/transactions"
)

type transactionService interface {
	Checkout(ctx context.Context, req transactions.CheckoutRequest) (float64, error)
}

type Handler struct {
	*gin.RouterGroup
	transctionSvc transactionService
}

func NewHandler(api *gin.RouterGroup, transctionSvc transactionService) *Handler {
	return &Handler{
		RouterGroup:   api,
		transctionSvc: transctionSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("transactions")
	route.POST("/checkouts", h.Checkout)
}
