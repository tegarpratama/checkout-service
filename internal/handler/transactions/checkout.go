package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tegarpratama/checkout-service/internal/model/transactions"
)

func (h *Handler) Checkout(c *gin.Context) {
	ctx := c.Request.Context()

	var request transactions.CheckoutRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	total, err := h.transctionSvc.Checkout(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "succesfully checkout",
		"total":   total,
	})
}
