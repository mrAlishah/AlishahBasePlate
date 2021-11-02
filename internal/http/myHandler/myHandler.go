package myHandler

import (
	"GolangTraining/internal/myService"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MyHandler struct {
	VersionInfo struct {
		GitCommit       string
		BuildTime       string
		ContainerName   string
		BinaryStartTime time.Time
	}
	HTTPServer *http.Server
	//01- assign myservice
	MyService myService.MyService
}

//01- CreateHandler Creates a new instance of REST handler
func CreateMyHandler(ss myService.MyService) *MyHandler {
	return &MyHandler{
		MyService: ss,
	}
}

//01- Start starts the http server
func (h *MyHandler) Start(ctx context.Context, port int32, r *gin.Engine) {
	const op = "http.rest.start"

	addr := fmt.Sprintf(":%d", port)

	h.HTTPServer = &http.Server{
		Addr:    addr,
		Handler: r,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	logrus.Infof("[OK] Starting HTTP REST Server on %s ", addr)
	err := h.HTTPServer.ListenAndServe()
	if err != http.ErrServerClosed {
		logrus.Fatal(errors.WithMessage(err, op))
	}
	// Code Reach Here after HTTP Server Shutdown!
	logrus.Info("[OK] HTTP REST Server is shutting down!")
}

//01- Stop handles the http server in graceful shutdown
func (h *MyHandler) Stop() {
	const op = "http.rest.stop"

	// Create an 5s timeout context or waiting for app to shutdown after 5 seconds
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelTimeout()

	h.HTTPServer.SetKeepAlivesEnabled(false)
	if err := h.HTTPServer.Shutdown(ctxTimeout); err != nil {
		logrus.Error(errors.WithMessage(err, op))
	}
	logrus.Info("HTTP REST Server graceful shutdown completed")
}
