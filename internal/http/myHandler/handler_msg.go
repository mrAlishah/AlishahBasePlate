package myHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *MyHandler) GetHelloHandler(c *gin.Context) {
	msg, err := h.MyService.Hello(c, "omid")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}
