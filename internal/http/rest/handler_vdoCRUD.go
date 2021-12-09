package rest

import (
	"AlishahBasePlate/internal/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
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

func (h *Handler) UploadVideo(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("videos/" + filename)
	if err != nil {
		logrus.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		logrus.Fatal(err)
	}
	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})

}
