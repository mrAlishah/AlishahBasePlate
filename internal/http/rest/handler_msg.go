package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetHelloHandler(c *gin.Context) {
	msg, err := h.Service.Hello(c, "omid")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}
