package rest

import (
	"GolangTraining/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (h *Handler) CreateVideo(c *gin.Context) {
	var request services.Video

	if err := c.ShouldBind(&request); err != nil {
		if vErr, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, GetFailResponseFromValidationErrors(vErr))
		}
		return
	}

	ctx := context.Background()
	user, err := h.Service.Create(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
