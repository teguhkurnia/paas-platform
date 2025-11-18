package repository

import (
	"gofiber-boilerplate/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (r *ProjectRepository) FindByIDWithServices(db *gorm.DB, entity *entity.Project, id uint) error {
	result := db.Preload("Services").First(&entity, id)
	if result.Error != nil {
		r.Logger.Errorf("Failed to find project by ID with services: %v", result.Error)
		return result.Error
	}
	return nil
}
