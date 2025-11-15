package repository

import (
	"gofiber-boilerplate/internal/entity"

	"github.com/sirupsen/logrus"
)

type ProjectRepository struct {
	Repository[entity.Project]
	Logger *logrus.Logger
}

func NewProjectRepository(logger *logrus.Logger) *ProjectRepository {
	return &ProjectRepository{
		Repository: Repository[entity.Project]{},
		Logger:     logger,
	}
}
