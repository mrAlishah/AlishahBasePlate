package services

import (
	"GolangTraining/internal/logger"
	"GolangTraining/internal/metrics"
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
)

//01- New Service and Add Service methods
type Service interface {
	Hello(ctx context.Context, name string) (string, error)

	Create(ctx context.Context, video Video) (Video, error)
}

//01-
type service struct {
	validate   *validator.Validate
	mysql      MySQLRepository
	redis      RedisRepository
	logger     *logger.StandardLogger
	prometheus *metrics.Prometheus
	config     *Config
}

//01- config
func CreateService(
	config *Config,
	logger *logger.StandardLogger,
	mysql MySQLRepository,
	redis RedisRepository,
	prometheus *metrics.Prometheus,
	validator *validator.Validate) Service {
	return &service{
		validate:   validator,
		redis:      redis,
		mysql:      mysql,
		logger:     logger,
		prometheus: prometheus,
		config:     config,
	}
}

//01- implimet service method or impliment into other file
func (s service) Hello(ctx context.Context, name string) (string, error) {
	//fmt.Printf("Hi %s", name)
	msg := fmt.Sprintf("Hi %s", name)
	return msg, nil

}
