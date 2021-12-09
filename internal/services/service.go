package services

import (
	"AlishahBasePlate/internal/logger"
	"AlishahBasePlate/internal/metrics"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
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
	postgres   PostgresRepository //02-
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
	postgres PostgresRepository, //02-
	redis RedisRepository,
	prometheus *metrics.Prometheus,
	validator *validator.Validate) Service {
	return &service{
		validate:   validator,
		redis:      redis,
		mysql:      mysql,
		postgres:   postgres,
		logger:     logger,
		prometheus: prometheus,
		config:     config,
	}
}

//01- implimet service method or impliment into other file
func (s service) Hello(ctx context.Context, name string) (string, error) {
	//fmt.Printf("Hi %s", name)
	msg := fmt.Sprintf("Hi %s", name)

	t := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				t.Stop()
				s.logger.Infof("updating movies context cancelled")
				return
			case <-t.C:
				s.printer()
			}
		}
	}()

	return msg, nil
}

func (s service) printer(){
	s.logger.Infof("Hiiii")
}
