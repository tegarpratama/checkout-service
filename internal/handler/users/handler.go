package users

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tegarpratama/checkout-service/internal/model/users"
)

type userService interface {
	CreateUser(ctx context.Context, req users.CreateUserRequest) (int64, error)
}

type Handler struct {
	*gin.RouterGroup
	userSvc userService
}

func NewHandler(api *gin.RouterGroup, userSvc userService) *Handler {
	return &Handler{
		RouterGroup: api,
		userSvc:     userSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("users")
	route.POST("/", h.CreateUser)
}
