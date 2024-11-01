package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tegarpratama/checkout-service/internal/model/users"
)

func (h *Handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var req users.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, err := h.userSvc.CreateUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := users.CreateUserResponse{
		ID:    userID,
		Email: req.Email,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created",
		"data":    response,
	})
}
