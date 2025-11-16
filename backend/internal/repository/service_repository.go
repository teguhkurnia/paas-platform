package repository

import (
	"gofiber-boilerplate/internal/entity"

	"github.com/sirupsen/logrus"
)

type ServiceRepository struct {
	Repository[entity.Service]
	Logger *logrus.Logger
}

func NewServiceRepository(logger *logrus.Logger) *ServiceRepository {
	return &ServiceRepository{
		Logger: logger,
	}
}
